package template

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	AccessKeyId     = ""
	AccessKeySecret = ""
)

func Test_templateClient_fetchTemplate(t *testing.T) {
	os.Setenv("ACCESS_KEY_ID", AccessKeyId)
	os.Setenv("ACCESS_KEY_SECRET", AccessKeySecret)
	client := NewClient("cd262474-dbb6-40fd-9c2a-3a4d7e53024f")
	filePath, _, err := client.FetchTemplate("HEAD")
	fmt.Println(filePath)
	assert.NoError(t, err)
	assert.FileExists(t, filePath)
}