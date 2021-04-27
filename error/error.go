package error

import (
	"com.poalim.bank.hackathon.login-fiber/global"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var (
	AccountNotExists = &fiber.Error{
		Code:    http.StatusUnprocessableEntity,
		Message: "account does not exist",
	}
	AccountAlreadyExist = &fiber.Error{
		Code:    http.StatusConflict,
		Message: "account already exists",
	}
	NoJwtInRequest = &fiber.Error{
		Code:    http.StatusBadRequest,
		Message: "no jwt sent",
	}
	IncorrectPassword = &fiber.Error{
		Code:    http.StatusBadRequest,
		Message: global.INCORRECT_PASSWORD,
	}
)
