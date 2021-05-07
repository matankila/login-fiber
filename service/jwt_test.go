package service

import (
	"com.poalim.bank.hackathon.login-fiber/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJwtWrapper(t *testing.T) {
	jw := jwtWrapper{}
	assert.Empty(t, jw.Issuer)
	assert.Empty(t, jw.SecretKey)
	assert.Equal(t, jw.ExpirationHours, int64(0))
}

func TestJwtClaim(t *testing.T) {
	jc := JwtClaim{}
	assert.Empty(t, jc.BankNumber)
	assert.Empty(t, jc.AccountId)
}

func TestNewJwtWrapper(t *testing.T) {
	jw := NewJwtWrapper()
	assert.Equal(t, jw.SecretKey, "ttlogin")
	assert.Equal(t, jw.Issuer, "ttlogin")
	assert.Equal(t, jw.ExpirationHours, int64(72))
}

func TestNewJwtWrapper2(t *testing.T) {
	jw1 := NewJwtWrapper()
	jw2 := NewJwtWrapper()
	assert.Equal(t, jw1, jw2)
}

func TestGenerateToken(t *testing.T) {
	jw := NewJwtWrapper()
	request := model.LoginRequest{
		BankNumber: "1234",
		AccountId:  "1234",
		Password:   "1234",
	}

	jwt, err := jw.GenerateToken(request)
	assert.Nil(t, err)
	assert.NotEmpty(t, jwt)
}

func TestValidateToken(t *testing.T) {
	jw := NewJwtWrapper()
	request := model.LoginRequest{
		BankNumber: "1234",
		AccountId:  "1234",
		Password:   "1234",
	}

	jwt, _ := jw.GenerateToken(request)
	c, err := jw.ValidateToken(jwt)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, c.AccountId, "1234")
	assert.Equal(t, c.BankNumber, "1234")
	assert.Equal(t, c.Issuer, "ttlogin")
}

func TestValidateToken2(t *testing.T) {
	jw := NewJwtWrapper()
	request := model.LoginRequest{
		BankNumber: "1234",
		AccountId:  "1234",
		Password:   "1234",
	}

	jwt, _ := jw.GenerateToken(request)
	jw2 := jwtWrapper{
		SecretKey:       "123",
		Issuer:          "tt",
		ExpirationHours: 1,
	}
	_, err := jw2.ValidateToken(jwt)
	assert.NotNil(t, err)
}

func TestValidateToken3(t *testing.T) {
	jw := jwtWrapper{
		SecretKey:       "123",
		Issuer:          "tt",
		ExpirationHours: 0,
	}

	request := model.LoginRequest{
		BankNumber: "1234",
		AccountId:  "1234",
		Password:   "1234",
	}
	jwt, _ := jw.GenerateToken(request)
	time.Sleep(1 * time.Second)
	_, err := jw.ValidateToken(jwt)
	assert.Error(t, err)
}
