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
)

// A thin wrapper around the "helm" command, adding logging and error translation.
type Cmd struct {
	HelmVer
	helmHome string
	WorkDir  string
	RepoType string
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
	cmd := exec.Command(c.binaryName, args...)
	cmd.Dir = c.WorkDir
	cmd.Env = os.Environ()
	if c.ociSupported {
		cmd.Env = append(cmd.Env,
			fmt.Sprint("HELM_EXPERIMENTAL_OCI=1"),
			fmt.Sprintf("XDG_CACHE_HOME=%s/cache", c.helmHome),
			fmt.Sprintf("XDG_CONFIG_HOME=%s/config", c.helmHome),
			fmt.Sprintf("XDG_DATA_HOME=%s/data", c.helmHome),
			fmt.Sprintf("HELM_HOME=%s", c.helmHome))
	} else {
		cmd.Env = append(cmd.Env,
			fmt.Sprintf("XDG_CACHE_HOME=%s/cache", c.helmHome),
			fmt.Sprintf("XDG_CONFIG_HOME=%s/config", c.helmHome),
			fmt.Sprintf("XDG_DATA_HOME=%s/data", c.helmHome),
			fmt.Sprintf("HELM_HOME=%s", c.helmHome))
	}

	return executil.RunWithRedactor(cmd, redactor)
}

func (c *Cmd) Init() (string, error) {
	if c.initSupported {
		return c.run("init", "--client-only", "--skip-refresh")
	}
	return "", nil
}

func (c *Cmd) Login(repo string, creds Creds) (string, error) {
	if c.ociSupported {
		args := []string{"registry", "login"}
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
		args := []string{"registry", "logout"}
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

	args := []string{"repo", "add"}

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

func (c *Cmd) Fetch(repoNamespace, repo, chartName, version, destination string, creds Creds) (string, error) {
	output := ""
	var err error
	if c.ociSupported {
		args := []string{"chart", "pull"}
		repoUrl := fmt.Sprintf(repo + "/" + repoNamespace + "/" + chartName + ":" + version)

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
		args = []string{"chart", "export", repoUrl, "--destination", destination}

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

		output, err = c.run(args...)
		if err != nil {
			return "", err
		}

		// tar helm chart
		cmd := exec.Command("tar", "-zcvf", repoNamespace+"-"+chartName+"-"+version+".tgz", chartName)
		cmd.Dir = destination
		_, err = executil.Run(cmd)
		if err != nil {
			return "", err
		}

		// remove helm chart
		cmd = exec.Command("rm", "-rf", chartName)
		cmd.Dir = destination
		_, err = executil.Run(cmd)
		if err != nil {
			return "", err
		}

	} else {
		args := []string{c.pullCommand, "--destination", destination}

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
		output, err = c.run(args...)
		if err != nil {
			return "", err
		}
	}

	return output, nil
}

func (c *Cmd) dependencyBuild() (string, error) {
	return c.run("dependency", "build")
}

func (c *Cmd) inspectValues(values string) (string, error) {
	return c.run(c.showCommand, "values", values)
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

	args := []string{"template", chartPath, c.templateNameArg, opts.Name}

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
