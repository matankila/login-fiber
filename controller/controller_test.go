package controller

import (
	"com.poalim.bank.hackathon.login-fiber/model"
	"com.poalim.bank.hackathon.login-fiber/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	res, err := hashPassword("matan")
	assert.Nil(t, err)
	assert.True(t, CheckPasswordHash("matan", res))
}

func TestLoginRequestToAccountData(t *testing.T) {
	request := model.LoginRequest{
		BankNumber: "1234",
		AccountId:  "1234",
		Password:   "1234",
	}

	res := loginRequestToAccountData(request)
	assert.Equal(t, res.Id, "1234-1234")
}

func TestRegisterRequestToAccountData(t *testing.T) {
	request := model.RegisterRequest{
		BankNumber: "1234",
		AccountId:  "1234",
		Password:   "1234",
	}

	res, err := registerRequestToAccountData(request)
	assert.Nil(t, err)
	assert.Equal(t, res.Id, "1234-1234")
	assert.True(t, CheckPasswordHash("1234", res.Password))
}

func TestHealth(t *testing.T) {
	err := Health()
	assert.Nil(t, err)
}

func TestValidate(t *testing.T) {
	err := Validate("x")
	assert.Error(t, err)
}

func TestValidate2(t *testing.T) {
	l := model.LoginRequest{
		BankNumber: "",
		AccountId:  "",
		Password:   "",
	}
	j := service.NewJwtWrapper()
	jwt, _ := j.GenerateToken(l)
	err := Validate(jwt)
	assert.Nil(t, err)
}

func TestLogin(t *testing.T) {
	r := model.LoginRequest{
		BankNumber: "m",
		AccountId:  "m",
		Password:   "m",
	}
	err := Login(r)
	assert.Error(t, err)
}

func TestRegister(t *testing.T) {
	r := model.RegisterRequest{
		BankNumber: "m",
		AccountId:  "m",
		Password:   "m",
	}
	err := Register(r)
	assert.Nil(t, err)
}

func TestRegister2(t *testing.T) {
	r := model.RegisterRequest{
		BankNumber: "m",
		AccountId:  "m",
		Password:   "m",
	}
	err := Register(r)
	assert.Error(t, err)
}
