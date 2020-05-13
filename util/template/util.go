package template

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"time"
)

const (
	// MetadataURL is metadata url
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
)

var (
	templateHashRegex = regexp.MustCompile("^[0-9a-f]{32}$")
)

// RoleAuth define STS Token Response
type RoleAuth struct {
	AccessKeyID     string
	AccessKeySecret string
	Expiration      time.Time
	SecurityToken   string
	LastUpdated     time.Time
	Code            string
}

// GetMetaData get metadata from ecs meta-server
func GetMetaData(resource string) (string, error) {
	resp, err := http.Get(MetadataURL + resource)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func IsRevisionHash(revision string) bool {
	return templateHashRegex.MatchString(revision)
}