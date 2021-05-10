package controller

import (
	"com.poalim.bank.hackathon.login-fiber/dao"
	"com.poalim.bank.hackathon.login-fiber/global"
	error_lib "com.poalim.bank.hackathon.login-fiber/global/error"
	"com.poalim.bank.hackathon.login-fiber/model"
	"com.poalim.bank.hackathon.login-fiber/service"
	"golang.org/x/crypto/bcrypt"
)

func Login(request model.LoginRequest) error {
	account := loginRequestToAccountData(request)

	c, _ := dao.New(global.URI)
	encPassword, err := c.Get(account)

	if err != nil {
		return err
	}
	if ok := CheckPasswordHash(account.Password, encPassword.(string)); !ok {
		return error_lib.IncorrectPassword
	}
	return nil
}

func Register(request model.RegisterRequest) error {
	account, err := registerRequestToAccountData(request)
	if err != nil {
		return err
	}

	c, _ := dao.New(global.URI)
	if _, err := c.Get(account); err == nil {
		return error_lib.AccountAlreadyExist
	}

	if err := c.Set(account); err != nil {
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
	c, _ := dao.New(global.URI)
	_, err := c.Ping()
	if err != nil {
		return err
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func loginRequestToAccountData(request model.LoginRequest) model.AccountData {
	return model.AccountData{
		Id:       request.BankNumber + "-" + request.AccountId,
		Password: request.Password,
	}
}

func registerRequestToAccountData(request model.RegisterRequest) (model.AccountData, error) {
	h, err := hashPassword(request.Password)
	if err != nil {
		return model.AccountData{}, err
	}

	return model.AccountData{
		Id:       request.BankNumber + "-" + request.AccountId,
		Password: h,
	}, nil
}
