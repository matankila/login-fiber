package controller

import (
	"com.poalim.bank.hackathon.login-fiber/dao"
	error_lib "com.poalim.bank.hackathon.login-fiber/global/error"
	"com.poalim.bank.hackathon.login-fiber/model"
	"com.poalim.bank.hackathon.login-fiber/service"
)

func Login(request model.LoginRequest) error {
	account := loginRequestToAccountData(request)
	encPassword, err := dao.GetAccount(account)

	if err != nil {
		return err
	}
	if encPassword == account.Password {
		return nil
	} else {
		return error_lib.IncorrectPassword
	}
}

func Register(request model.RegisterRequest) error {
	account := registerRequestToAccountData(request)
	if _, err := dao.GetAccount(account); err == nil {
		return error_lib.AccountAlreadyExist
	}

	if err := dao.SetAccount(account); err != nil {
		return err
	}
	return nil
}

func Validate(jwt string) error {
	j := service.NewJwtWrapper()
	_, err := j.ValidateToken(jwt)
	if err != nil {
		return err
	}
	return nil
}

func Health() error {
	_, err := dao.Ping()
	if err != nil {
		return err
	}
	return nil
}

func hashPassword(password string) string {
	return "tom-" + password
}

func loginRequestToAccountData(request model.LoginRequest) model.AccountData {
	return model.AccountData{
		Id:       request.BankNumber + "-" + request.AccountId,
		Password: hashPassword(request.Password),
	}
}

func registerRequestToAccountData(request model.RegisterRequest) model.AccountData {
	return model.AccountData{
		Id:       request.BankNumber + "-" + request.AccountId,
		Password: hashPassword(request.Password),
	}
}
