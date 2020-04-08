package helm

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"

	"github.com/argoproj/argo-cd/util"
	executil "github.com/argoproj/argo-cd/util/exec"
	"strings"
)

// A thin wrapper around the "helm" command, adding logging and error translation.
type Cmd struct {
	HelmVer
	helmHome string
	WorkDir  string
}

func NewCmd(workDir string) (*Cmd, error) {
	helmVersion, err := getHelmVersion(workDir)
	if err != nil {
		return nil, err
	}

	return NewCmdWithVersion(workDir, *helmVersion)
}

func NewCmdWithVersion(workDir string, version HelmVer) (*Cmd, error) {
	tmpDir, err := ioutil.TempDir("", "helm")
	if err != nil {
		return nil, err
	}
	return &Cmd{WorkDir: workDir, helmHome: tmpDir, HelmVer: version}, err
}

var redactor = func(text string) string {
	return regexp.MustCompile("(--username|--password) [^ ]*").ReplaceAllString(text, "$1 ******")
}

func (c Cmd) run(args ...string) (string, error) {
	argsStr := strings.Join(args, " ")
	cmd := exec.Command("sh", "-c", argsStr)
	cmd.Dir = c.WorkDir
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env,
		fmt.Sprintf("XDG_CACHE_HOME=%s/cache", c.helmHome),
		fmt.Sprintf("XDG_CONFIG_HOME=%s/config", c.helmHome),
		fmt.Sprintf("XDG_DATA_HOME=%s/data", c.helmHome),
		fmt.Sprintf("HELM_HOME=%s", c.helmHome))
	return executil.RunWithRedactor(cmd, redactor)
}

func (c *Cmd) Init() (string, error) {
	if c.initSupported {
		return c.run(c.binaryName, "init", "--client-only", "--skip-refresh")
	}
	return "", nil
}

func (c *Cmd) Login(repo string, creds Creds) (string, error) {
	if c.ociSupported {
		args := []string{HelmOCIEnv, c.binaryName, "registry", "login"}
		args = append(args, repo)

		if creds.Username != "" {
			args = append(args, "--username", creds.Username)
		}

		if creds.Password != "" {
			args = append(args, "--password", creds.Password)
		}

		if creds.CAPath != "" {
			args = append(args, "--ca-file", creds.CAPath)
		}
		if len(creds.CertData) > 0 {
			filePath, closer, err := writeToTmp(creds.CertData)
			if err != nil {
				return "", err
			}
			defer util.Close(closer)
			args = append(args, "--cert-file", filePath)
		}
		if len(creds.KeyData) > 0 {
			filePath, closer, err := writeToTmp(creds.KeyData)
			if err != nil {
				return "", err
			}
			defer util.Close(closer)
			args = append(args, "--key-file", filePath)
		}

		return c.run(args...)
	}
	return "", nil
}

func (c *Cmd) Logout(repo string, creds Creds) (string, error) {
	if c.ociSupported {
		args := []string{HelmOCIEnv, c.binaryName, "registry", "logout"}
		args = append(args, repo)

		if creds.CAPath != "" {
			args = append(args, "--ca-file", creds.CAPath)
		}
		if len(creds.CertData) > 0 {
			filePath, closer, err := writeToTmp(creds.CertData)
			if err != nil {
				return "", err
			}
			defer util.Close(closer)
			args = append(args, "--cert-file", filePath)
		}
		if len(creds.KeyData) > 0 {
			filePath, closer, err := writeToTmp(creds.KeyData)
			if err != nil {
				return "", err
			}
			defer util.Close(closer)
			args = append(args, "--key-file", filePath)
		}

		return c.run(args...)
	}
	return "", nil
}

func (c *Cmd) RepoAdd(name string, url string, opts Creds) (string, error) {
	tmp, err := ioutil.TempDir("", "helm")
	if err != nil {
		return "", err
	}
	defer func() { _ = os.RemoveAll(tmp) }()

	args := []string{c.binaryName, "repo", "add"}

	if opts.Username != "" {
		args = append(args, "--username", opts.Username)
	}

	if opts.Password != "" {
		args = append(args, "--password", opts.Password)
	}

	if opts.CAPath != "" {
		args = append(args, "--ca-file", opts.CAPath)
	}

	if len(opts.CertData) > 0 {
		certFile, err := ioutil.TempFile("", "helm")
		if err != nil {
			return "", err
		}
		_, err = certFile.Write(opts.CertData)
		if err != nil {
			return "", err
		}
		args = append(args, "--cert-file", certFile.Name())
	}

	if len(opts.KeyData) > 0 {
		keyFile, err := ioutil.TempFile("", "helm")
		if err != nil {
			return "", err
		}
		_, err = keyFile.Write(opts.KeyData)
		if err != nil {
			return "", err
		}
		args = append(args, "--key-file", keyFile.Name())
	}

	args = append(args, name, url)

	return c.run(args...)
}

func writeToTmp(data []byte) (string, io.Closer, error) {
	file, err := ioutil.TempFile("", "")
	if err != nil {
		return "", nil, err
	}
	err = ioutil.WriteFile(file.Name(), data, 0644)
	if err != nil {
		_ = os.RemoveAll(file.Name())
		return "", nil, err
	}
	return file.Name(), util.NewCloser(func() error {
		return os.RemoveAll(file.Name())
	}), nil
}

func (c *Cmd) Fetch(repo, chartName, version, destination string, creds Creds) (string, error) {
	args := []string{}
	if c.ociSupported {
		args = []string{HelmOCIEnv, c.binaryName, "chart", "pull"}
		repoUrl := ""
		if version != "" {
			repoUrl = fmt.Sprintf(repo+"/"+chartName+":"+version)
		} else {
			repoUrl = fmt.Sprintf(repo+"/"+chartName)
		}

		args = append(args, repoUrl)

		if creds.CAPath != "" {
			args = append(args, "--ca-file", creds.CAPath)
		}
		if len(creds.CertData) > 0 {
			filePath, closer, err := writeToTmp(creds.CertData)
			if err != nil {
				return "", err
			}
			defer util.Close(closer)
			args = append(args, "--cert-file", filePath)
		}
		if len(creds.KeyData) > 0 {
			filePath, closer, err := writeToTmp(creds.KeyData)
			if err != nil {
				return "", err
			}
			defer util.Close(closer)
			args = append(args, "--key-file", filePath)
		}

		_, err := c.run(args...)
		if err != nil {
			return "", err
		}

		// helm chart export
		args = []string{HelmOCIEnv, c.binaryName, "chart", "export", repoUrl, "--destination", destination}

		if creds.CAPath != "" {
			args = append(args, "--ca-file", creds.CAPath)
		}
		if len(creds.CertData) > 0 {
			filePath, closer, err := writeToTmp(creds.CertData)
			if err != nil {
				return "", err
			}
			defer util.Close(closer)
			args = append(args, "--cert-file", filePath)
		}
		if len(creds.KeyData) > 0 {
			filePath, closer, err := writeToTmp(creds.KeyData)
			if err != nil {
				return "", err
			}
			defer util.Close(closer)
			args = append(args, "--key-file", filePath)
		}

	} else {
		args = []string{c.binaryName, c.pullCommand, "--destination", destination}

		if version != "" {
			args = append(args, "--version", version)
		}
		if creds.Username != "" {
			args = append(args, "--username", creds.Username)
		}
		if creds.Password != "" {
			args = append(args, "--password", creds.Password)
		}
		if creds.CAPath != "" {
			args = append(args, "--ca-file", creds.CAPath)
		}
		if len(creds.CertData) > 0 {
			filePath, closer, err := writeToTmp(creds.CertData)
			if err != nil {
				return "", err
			}
			defer util.Close(closer)
			args = append(args, "--cert-file", filePath)
		}
		if len(creds.KeyData) > 0 {
			filePath, closer, err := writeToTmp(creds.KeyData)
			if err != nil {
				return "", err
			}
			defer util.Close(closer)
			args = append(args, "--key-file", filePath)
		}

		args = append(args, "--repo", repo, chartName)
	}

	return c.run(args...)
}

func (c *Cmd) dependencyBuild() (string, error) {
	return c.run(c.binaryName, "dependency", "build")
}

func (c *Cmd) inspectValues(values string) (string, error) {
	return c.run(c.binaryName, c.showCommand, "values", values)
}

type TemplateOpts struct {
	Name        string
	Namespace   string
	KubeVersion string
	APIVersions []string
	Set         map[string]string
	SetString   map[string]string
	SetFile     map[string]string
	Values      []string
}

var (
	re = regexp.MustCompile(`([^\\]),`)
)

func cleanSetParameters(val string) string {
	return re.ReplaceAllString(val, `$1\,`)
}

func (c *Cmd) template(chartPath string, opts *TemplateOpts) (string, error) {
	if c.HelmVer.getPostTemplateCallback != nil {
		if callback, err := c.HelmVer.getPostTemplateCallback(filepath.Clean(path.Join(c.WorkDir, chartPath))); err == nil {
			defer callback()
		} else {
			return "", err
		}
	}

	args := []string{c.binaryName, "template", chartPath, c.templateNameArg, opts.Name}

	if opts.Namespace != "" {
		args = append(args, "--namespace", opts.Namespace)
	}
	if opts.KubeVersion != "" && c.kubeVersionSupported {
		args = append(args, "--kube-version", opts.KubeVersion)
	}
	for key, val := range opts.Set {
		args = append(args, "--set", key+"="+cleanSetParameters(val))
	}
	for key, val := range opts.SetString {
		args = append(args, "--set-string", key+"="+cleanSetParameters(val))
	}
	for key, val := range opts.SetFile {
		args = append(args, "--set-file", key+"="+cleanSetParameters(val))
	}
	for _, val := range opts.Values {
		args = append(args, "--values", val)
	}
	for _, v := range opts.APIVersions {
		args = append(args, "--api-versions", v)
	}

	return c.run(args...)
}

func (c *Cmd) Close() {
	_ = os.RemoveAll(c.helmHome)
}
