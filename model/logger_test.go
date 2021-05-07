package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestInfo(t *testing.T) {
	r := RequestInfo{}
	assert.Empty(t, r.Url)
	assert.Empty(t, r.Ip)
	assert.Empty(t, r.Method)
	assert.Empty(t, r.UID)
}
