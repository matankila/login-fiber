package controller

import (
	"com.poalim.bank.hackathon.login-fiber/dao"
	error_lib "com.poalim.bank.hackathon.login-fiber/error"
	"com.poalim.bank.hackathon.login-fiber/global"
	"com.poalim.bank.hackathon.login-fiber/model"
)

func Login(request model.LoginRequest) (model.LoginResponse, error) {
	response := model.LoginResponse{}
	account := loginRequestToAccountData(request)
	encPassword, err := dao.GetAccount(account)

	if err != nil {
		response.Ok = false
		response.Message = err.Error()
		return response, err
	}

	if encPassword == account.Password {
		response.Ok = true
		response.Message = global.LOGIN_RESPONSE
		return response, nil
	} else {
		response.Ok = false
		response.Message = global.INCORRECT_PASSWORD
		return response, error_lib.IncorrectPassword
	}
}

func Register(request model.RegisterRequest) (model.RegisterResponse, error) {
	response := model.RegisterResponse{}
	account := registerRequestToAccountData(request)
	if _, err := dao.GetAccount(account); err == nil {
		return response, error_lib.AccountAlreadyExist
	}
	err := dao.SetAccount(account)

	if err != nil {
		response.Ok = false
		response.Message = err.Error()
		return response, err
	}

	response.Ok = true
	response.Message = global.REGISTER_RESPONSE
	return response, nil

}

func Validate(jwt string) (model.ValidateResponse, error) {
	response := model.ValidateResponse{}
	_, err := ValidateJwt(jwt)
	if err != nil {
		response.Ok = false
		response.Message = err.Error()
		return response, err
	}
	response.Ok = true
	response.Message = global.VALIDATE_RESPONSE
	return response, nil
}

func Health() (model.HealthResponse, error) {
	response := model.HealthResponse{}
	_, err := dao.Ping()
	if err != nil {
		response.Ok = false
		response.Message = err.Error()
		return response, err
	}

	response.Ok = true
	response.Message = global.HEALTH_RESPONSE
	return response, nil
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
