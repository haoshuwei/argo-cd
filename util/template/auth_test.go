package template

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetManagedToken(t *testing.T) {
	accessKeyID, accessSecret, accessToken := GetManagedToken()
	fmt.Println(accessKeyID, accessSecret, accessToken)
	assert.NotEmpty(t, accessToken)
}
