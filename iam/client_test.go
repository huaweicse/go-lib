package iam_test

import (
	"github.com/huaweicse/iam"
	"github.com/stretchr/testify/assert"
	"testing"
)

var c *iam.Client

func TestNew(t *testing.T) {
	var err error
	c, err = iam.NewClient(iam.Options{
		Endpoint:  "https://iam.cn-north-1.myhuaweicloud.com",
		AccessKey: "xxx",
		SecretKey: "xxx",
		Project:   "default",
	})
	assert.NoError(t, err)

}
func TestClient_GetToken(t *testing.T) {
	s, err := c.GetToken("xxx", "xxx..", "xxx")
	assert.NoError(t, err)
	t.Log(s)
}
