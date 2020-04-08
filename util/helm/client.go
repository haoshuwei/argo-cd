package helm

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Masterminds/semver"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/argoproj/argo-cd/util"
	"os/exec"
	executil "github.com/argoproj/argo-cd/util/exec"
)

var (
	globalLock = util.NewKeyLock()
)

type Creds struct {
	Username string
	Password string
	CAPath   string
	CertData []byte
	KeyData  []byte
}

type Client interface {
	CleanChartCache(chart string, version *semver.Version) error
	ExtractChart(chart string, version *semver.Version) (string, util.Closer, error)
	GetIndex() (*Index, error)
}

func NewClient(repoType string, repoURL string, creds Creds) Client {
	return NewClientWithLock(repoType, repoURL, creds, globalLock)
}

func NewClientWithLock(repoType string, repoURL string, creds Creds, repoLock *util.KeyLock) Client {
	return &nativeHelmChart{
		repoType: repoType,
		repoURL:  repoURL,
		creds:    creds,
		repoPath: filepath.Join(os.TempDir(), strings.Replace(repoURL, "/", "_", -1)),
		repoLock: repoLock,
	}
}

type nativeHelmChart struct {
	repoType string
	repoPath string
	repoURL  string
	creds    Creds
	repoLock *util.KeyLock
}

func fileExist(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (c *nativeHelmChart) ensureHelmChartRepoPath() error {
	c.repoLock.Lock(c.repoPath)
	defer c.repoLock.Unlock(c.repoPath)

	err := os.Mkdir(c.repoPath, 0700)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func (c *nativeHelmChart) CleanChartCache(chart string, version *semver.Version) error {
	return os.RemoveAll(c.getChartPath(c.repoType, chart, version))
}

func (c *nativeHelmChart) ExtractChart(chart string, version *semver.Version) (string, util.Closer, error) {
	err := c.ensureHelmChartRepoPath()
	if err != nil {
		return "", nil, err
	}
	chartPath := c.getChartPath(c.repoType, chart, version)

	c.repoLock.Lock(chartPath)
	defer c.repoLock.Unlock(chartPath)

	exists, err := fileExist(chartPath)
	if err != nil {
		return "", nil, err
	}
	if !exists {
		// always use Helm V3 since we don't have chart content to determine correct Helm version
		helmCmd := &Cmd{}
		var err error
		if c.repoType == "helm-oci" {
			helmCmd, err = NewCmdWithVersion(c.repoPath, HelmOCI)
		} else {
			helmCmd, err = NewCmdWithVersion(c.repoPath, HelmV3)
		}
		if err != nil {
			return "", nil, err
		}
		defer helmCmd.Close()

		_, err = helmCmd.Init()
		if err != nil {
			return "", nil, err
		}

		_, err = helmCmd.Login(c.repoURL,c.creds)
		if err != nil {
			return "", nil, err
		}
		defer func() {
			_, _ = helmCmd.Logout(c.repoURL,c.creds)
		}()

		// (1) because `helm fetch` downloads an arbitrary file name, we download to an empty temp directory
		tempDest, err := ioutil.TempDir("", "helm")
		if err != nil {
			return "", nil, err
		}
		defer func() { _ = os.RemoveAll(tempDest) }()

		_, err = helmCmd.Fetch(c.repoURL, chart, version.String(), tempDest, c.creds)
		if err != nil {
			return "", nil, err
		}
		// (2) then we assume that the only file downloaded into the directory is the tgz file
		// and we move that to where we want it
		infos, err := ioutil.ReadDir(tempDest)
		if err != nil {
			return "", nil, err
		}
		if len(infos) != 1 {
			return "", nil, fmt.Errorf("expected 1 file, found %v", len(infos))
		}
		err = os.Rename(filepath.Join(tempDest, infos[0].Name()), chartPath)
		if err != nil {
			return "", nil, err
		}
	}
	// untar helm chart into throw away temp directory which should be deleted as soon as no longer needed
	tempDir, err := ioutil.TempDir("", "helm")
	if err != nil {
		return "", nil, err
	}
	chartName := ""
	if c.repoType == "helm-oci" {
		chartName = strings.Split(chart, "/")[1]
		cmd := exec.Command("cp", "-r", chartPath)
		cmd.Dir = tempDir
		_, err = executil.Run(cmd)
		if err != nil {
			_ = os.RemoveAll(tempDir)
			return "", nil, err
		}
	} else {
		chartName = chart
		cmd := exec.Command("tar", "-zxvf", chartPath)
		cmd.Dir = tempDir
		_, err = executil.Run(cmd)
		if err != nil {
			_ = os.RemoveAll(tempDir)
			return "", nil, err
		}
	}

	return path.Join(tempDir, chartName), util.NewCloser(func() error {
		return os.RemoveAll(tempDir)
	}), nil
}

func (c *nativeHelmChart) GetIndex() (*Index, error) {
	start := time.Now()

	data, err := c.loadRepoIndex()
	if err != nil {
		return nil, err
	}

	index := &Index{}
	err = yaml.NewDecoder(bytes.NewBuffer(data)).Decode(index)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{"seconds": time.Since(start).Seconds()}).Info("took to get index")

	return index, nil
}

func (c *nativeHelmChart) loadRepoIndex() ([]byte, error) {
	repoURL, err := url.Parse(c.repoURL)
	if err != nil {
		return nil, err
	}
	repoURL.Path = path.Join(repoURL.Path, "index.yaml")

	req, err := http.NewRequest("GET", repoURL.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.creds.Username != "" || c.creds.Password != "" {
		// only basic supported
		req.SetBasicAuth(c.creds.Username, c.creds.Password)
	}

	tlsConf, err := newTLSConfig(c.creds)
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: tlsConf,
	}
	client := http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to get index: " + resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}

func newTLSConfig(creds Creds) (*tls.Config, error) {
	tlsConfig := &tls.Config{InsecureSkipVerify: false}

	if creds.CAPath != "" {
		caData, err := ioutil.ReadFile(creds.CAPath)
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caData)
		tlsConfig.RootCAs = caCertPool
	}

	// If a client cert & key is provided then configure TLS config accordingly.
	if len(creds.CertData) > 0 && len(creds.KeyData) > 0 {
		cert, err := tls.X509KeyPair(creds.CertData, creds.KeyData)
		if err != nil {
			return nil, err
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
	}
	tlsConfig.BuildNameToCertificate()

	return tlsConfig, nil
}

func (c *nativeHelmChart) getChartPath(repoType string, chart string, version *semver.Version) string {
	if repoType == "helm-oci" {
		return path.Join(c.repoPath, fmt.Sprintf("%s-%v", chart, version))
	}
	return path.Join(c.repoPath, fmt.Sprintf("%s-%v.tgz", chart, version))
}
