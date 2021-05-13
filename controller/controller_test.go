package controller

import (
	"com.poalim.bank.hackathon.login-fiber/mock_dao"
	"com.poalim.bank.hackathon.login-fiber/mock_service"
	"com.poalim.bank.hackathon.login-fiber/model"
	"com.poalim.bank.hackathon.login-fiber/service"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoginRequestToAccountData(t *testing.T) {
	request := model.LoginRequest{
		BankNumber: "1234",
		AccountId:  "1234",
		Password:   "1234",
	}

	res := loginRequestToAccountData(request)
	assert.Equal(t, res.Id, "1234-1234")
	assert.Equal(t, res.Password, "1234")
}

func TestRegisterRequestToAccountData(t *testing.T) {
	request := model.RegisterRequest{
		BankNumber: "1234",
		AccountId:  "1234",
		Password:   "1234",
	}

	res := registerRequestToAccountData(request)
	assert.Equal(t, res.Id, "1234-1234")
	assert.Equal(t, res.Password, "1234")
}

func TestHealth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db := mock_dao.NewMockDB(ctrl)
	hash := mock_service.NewMockBycrypt(ctrl)
	db.EXPECT().Ping().Return(true, nil)
	c := NewController(db, hash)
	err := c.Health()
	assert.Nil(t, err)
}

func TestValidate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db := mock_dao.NewMockDB(ctrl)
	hash := mock_service.NewMockBycrypt(ctrl)
	c := NewController(db, hash)
	err := c.Validate("x")
	assert.Error(t, err)
}

func TestValidate2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db := mock_dao.NewMockDB(ctrl)
	hash := mock_service.NewMockBycrypt(ctrl)
	c := NewController(db, hash)
	l := model.LoginRequest{
		BankNumber: "",
		AccountId:  "",
		Password:   "",
	}
	j := service.NewJwtWrapper()
	jwt, _ := j.GenerateToken(l)
	err := c.Validate(jwt)
	assert.Nil(t, err)
}

func TestLogin(t *testing.T) {
	r := model.LoginRequest{
		BankNumber: "m",
		AccountId:  "m",
		Password:   "m",
	}
	account := loginRequestToAccountData(r)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	hash := mock_service.NewMockBycrypt(ctrl)
	db := mock_dao.NewMockDB(ctrl)
	db.EXPECT().Get(account).Return(nil, errors.New("not found"))
	c := NewController(db, hash)
	err := c.Login(r)
	assert.Error(t, err)
}

func TestLogin2(t *testing.T) {
	r := model.LoginRequest{
		BankNumber: "m",
		AccountId:  "m",
		Password:   "m",
	}
	account := loginRequestToAccountData(r)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	hash := mock_service.NewMockBycrypt(ctrl)
	db := mock_dao.NewMockDB(ctrl)
	db.EXPECT().Get(account).Return("***", nil)
	hash.EXPECT().CheckPasswordHash(account.Password, "***").Return(true)
	c := NewController(db, hash)
	err := c.Login(r)
	assert.Nil(t, err)
}

func TestRegister(t *testing.T) {
	r := model.RegisterRequest{
		BankNumber: "m",
		AccountId:  "m",
		Password:   "m",
	}
	account := registerRequestToAccountData(r)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	hash := mock_service.NewMockBycrypt(ctrl)
	db := mock_dao.NewMockDB(ctrl)
	hash.EXPECT().HashPassword(account.Password).Return("***", nil)
	account.Password = "***"
	db.EXPECT().Get(account).Return(nil, errors.New("not found"))
	db.EXPECT().Set(account).Return(nil)
	c := NewController(db, hash)
	err := c.Register(r)
	assert.Nil(t, err)
}

func TestRegister2(t *testing.T) {

	r := model.RegisterRequest{
		BankNumber: "m",
		AccountId:  "m",
		Password:   "m",
	}
	account := registerRequestToAccountData(r)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	hash := mock_service.NewMockBycrypt(ctrl)
	db := mock_dao.NewMockDB(ctrl)
	hash.EXPECT().HashPassword(account.Password).Return("***", nil)
	account.Password = "***"
	db.EXPECT().Get(account).Return(nil, nil)
	c := NewController(db, hash)
	err := c.Register(r)
	assert.Error(t, err)
}
