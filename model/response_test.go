package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoginResponse(t *testing.T) {
	r := LoginResponse{}
	assert.Empty(t, r.Message)
	assert.Empty(t, r.Jwt)
	assert.False(t, r.Ok)
}

func TestRegisterResponse(t *testing.T) {
	r := RegisterResponse{}
	assert.Empty(t, r.Message)
	assert.False(t, r.Ok)
}

func TestValidateResponse(t *testing.T) {
	r := ValidateResponse{}
	assert.Empty(t, r.Message)
	assert.False(t, r.Ok)
}

func TestHealthResponse(t *testing.T) {
	r := HealthResponse{}
	assert.Empty(t, r.Message)
	assert.False(t, r.Ok)
}

func TestErrorResponse(t *testing.T) {
	r := ErrorResponse{}
	assert.Empty(t, r.Message)
	assert.False(t, r.Ok)
}
