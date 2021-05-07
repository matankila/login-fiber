package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoginRequest(t *testing.T) {
	r := LoginRequest{}
	assert.Empty(t, r.Password)
	assert.Empty(t, r.AccountId)
	assert.Empty(t, r.BankNumber)
}

func TestRegisterRequest(t *testing.T) {
	r := RegisterRequest{}
	assert.Empty(t, r.Password)
	assert.Empty(t, r.AccountId)
	assert.Empty(t, r.BankNumber)
}
