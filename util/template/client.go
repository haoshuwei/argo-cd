package template

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/argoproj/argo-cd/util"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
)

type Client interface {
	FetchTemplate(revision string) (string, util.Closer, error)
	ResolveHeadRevision() (string, error)
	GetRevisionDetail(string) (*template, error)
}

type templateClient struct {
	id           string
	aliyunClient *sdk.Client
}

type template struct {
	Id          string `json:"template_with_hist_id"`
	Name        string `json:"name"`
	Template    string `json:"template"`
	Revision    string `json:"template_hash_code_version,omitempty"`
	Description string `json:"description,omitempty"`
	Updated     string `json:"updated"`
}

func NewClient(id string) Client {
	return &templateClient{id: id}
}

func (c *templateClient) FetchTemplate(revision string) (string, util.Closer, error) {
	t, err := c.fetchTemplate(revision)
	if err != nil {
		return "", nil, err
	}

	filePath, err := c.getFilePath(revision)
	err = ioutil.WriteFile(filePath, []byte(t.Template), 0700)
	if err != nil {
		return "", nil, err
	}
	return filePath, util.NewCloser(func() error {
		return os.RemoveAll(filePath)
	}), nil
}

func (c *templateClient) ResolveHeadRevision() (string, error) {
	t, err := c.fetchTemplate("HEAD")
	if err != nil {
		return "", err
	}

	return t.Revision, nil
}

func (c *templateClient) GetRevisionDetail(revision string) (*template, error) {
	if !IsRevisionHash(revision) {
		return nil, fmt.Errorf("invalid template revision")
	}

	t, err := c.fetchTemplate(revision)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *templateClient) fetchTemplate(revision string) (*template, error) {
	err := c.newAliyunClient()
	if err != nil {
		return nil, err
	}

	request := requests.NewCommonRequest()
	request.Method = "GET"
	request.Product = "CS"
	request.Domain = "cs.aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = fmt.Sprintf("/templates/%s", c.id)
	request.TransToAcsRequest()

	response, err := c.aliyunClient.ProcessCommonRequest(request)
	if err != nil {
		return nil, err
	}

	body := response.GetHttpContentString()
	var templates []template
	err = json.Unmarshal([]byte(body), &templates)
	if err != nil {
		return nil, err
	}

	if len(templates) == 0 {
		return nil, fmt.Errorf("versions not found for template %s", c.id)
	}

	log.Infof("Got template from ack: %s %s", templates[0].Id, templates[0].Revision)

	if revision == "HEAD" {
		return &templates[0], nil
	}

	for _, t := range templates {
		if t.Revision == revision {
			return &t, nil
		}
	}

	return nil, fmt.Errorf("versions not found for template %s", c.id)
}

func (c *templateClient) newAliyunClient() error {
	accessKeyId, accessKeySecret, accessToken := GetDefaultAK()
	if accessToken == "" {
		client, err := sdk.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessKeySecret)
		if err != nil {
			return err
		}
		c.aliyunClient = client
	} else {
		client, err := sdk.NewClientWithStsToken("cn-hangzhou", accessKeyId, accessKeySecret, accessToken)
		if err != nil {
			return err
		}
		c.aliyunClient = client
	}

	return nil
}

func (c *templateClient) getFilePath(revision string) (string, error) {
	fileDir := os.TempDir()
	filePath := path.Join(fileDir, fmt.Sprintf("%s-%v.yaml", c.id, revision))
	err := ensureFileDir(fileDir)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func ensureFileDir(dir string) error {
	err := os.MkdirAll(dir, 0700)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}
