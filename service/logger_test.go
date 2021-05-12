package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitFactory(t *testing.T) {
	c := InitLoggerFactory()
	assert.NotNil(t, c)
}

func TestGetLogger(t *testing.T) {
	InitLoggerFactory()
	l := GetLogger(Default)
	assert.NotNil(t, l)
}

func TestDefaultLogger_String(t *testing.T) {
	s := Default.String()
	assert.Equal(t, s, "default")
}

func TestHealthLogger_String(t *testing.T) {
	s := Health.String()
	assert.Equal(t, s, "health")
}
