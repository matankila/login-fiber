package controller

import (
	"com.poalim.bank.hackathon.login-fiber/dao"
	error_lib "com.poalim.bank.hackathon.login-fiber/global/error"
	"com.poalim.bank.hackathon.login-fiber/model"
	"com.poalim.bank.hackathon.login-fiber/service"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	db dao.DB
}

func NewContoller(db dao.DB) Controller {
	return Controller{
		db: db,
	}
}

func (c Controller) Login(request model.LoginRequest) error {
	account := loginRequestToAccountData(request)

	encPassword, err := c.db.Get(account)
	if err != nil {
		return err
	}

	if ok := checkPasswordHash(account.Password, encPassword.(string)); !ok {
		return error_lib.IncorrectPassword
	}

	return nil
}

func (c Controller) Register(request model.RegisterRequest) error {
	account, err := registerRequestToAccountData(request)
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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
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
