package controller

import (
	"com.poalim.bank.hackathon.login-fiber/dao"
	error_lib "com.poalim.bank.hackathon.login-fiber/global/error"
	"com.poalim.bank.hackathon.login-fiber/model"
	"com.poalim.bank.hackathon.login-fiber/service"
)

type Controller struct {
	db   dao.DB
	hash service.Bycrypt
}

func NewController(db dao.DB, hash service.Bycrypt) Controller {
	return Controller{
		db:   db,
		hash: hash,
	}
}

func (c Controller) Login(request model.LoginRequest) error {
	account := loginRequestToAccountData(request)

	encPassword, err := c.db.Get(account)
	if err != nil {
		return err
	}

	if ok := c.hash.CheckPasswordHash(account.Password, encPassword.(string)); !ok {
		return error_lib.IncorrectPassword
	}

	return nil
}

func (c Controller) Register(request model.RegisterRequest) error {
	var err error
	account := registerRequestToAccountData(request)

	account.Password, err = c.hash.HashPassword(account.Password)
	if err != nil {
		return err
	}

	if _, err := c.db.Get(account); err == nil {
		return error_lib.AccountAlreadyExist
	}

	if err := c.db.Set(account); err != nil {
		return err
	}

	return nil
}

func (c Controller) Validate(jwt string) error {
	j := service.NewJwtWrapper()
	_, err := j.ValidateToken(jwt)
	if err != nil {
		return err
	}
	return nil
}

func (c Controller) Health() error {
	_, err := c.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func loginRequestToAccountData(request model.LoginRequest) model.AccountData {
	return model.AccountData{
		Id:       request.BankNumber + "-" + request.AccountId,
		Password: request.Password,
	}
}

func registerRequestToAccountData(request model.RegisterRequest) model.AccountData {
	return model.AccountData{
		Id:       request.BankNumber + "-" + request.AccountId,
		Password: request.Password,
	}
}
