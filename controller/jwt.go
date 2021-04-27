package controller

import "com.poalim.bank.hackathon.login-fiber/model"

func CreateJwt(request model.LoginRequest) (string, error) {
	return "fake jwt", nil

}
func ValidateJwt(jwt string) (bool, error) {
	return true, nil
}
