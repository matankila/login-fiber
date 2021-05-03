package api

import (
	"com.poalim.bank.hackathon.login-fiber/controller"
	"com.poalim.bank.hackathon.login-fiber/global"
	"com.poalim.bank.hackathon.login-fiber/model"
	error_lib "com.poalim.bank.hackathon.login-fiber/model/error"
	"com.poalim.bank.hackathon.login-fiber/service"
	"github.com/gofiber/fiber/v2"
)

type handler func(ctx *fiber.Ctx) error

type Controller struct {
	login    handler
	register handler
	validate handler
	health   handler
}

func InitController() Controller {
	c := Controller{}
	c.login = LoginController
	c.register = RegisterController
	c.validate = ValidateController
	c.health = HealthController
	return c
}

// LoginController godoc
// @Summary login to app
// @Description login to app
// @ID login-to-app
// @Accept  json
// @Produce  json
// @Tags Login
// @Param account body model.LoginRequest true "login account"
// @Success 200 {object} model.LoginResponse
// @Failure 422 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /v1/login [post]
func LoginController(c *fiber.Ctx) error {
	request := model.LoginRequest{}
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	err := controller.Login(request)
	if err != nil {
		return err
	}
	j := service.NewJwtWrapper()
	jwt, err := j.GenerateToken(request)
	if err != nil {
		return err
	}

	return c.JSON(model.LoginResponse{
		Ok:      true,
		Message: global.LOGIN_RESPONSE,
		Jwt:     jwt,
	})
}

// RegisterController godoc
// @Summary register to app
// @Description register to app
// @ID register-to-app
// @Accept  json
// @Produce  json
// @Tags Register
// @Param account body model.RegisterRequest true "register account"
// @Success 200 {object} model.RegisterResponse
// @Failure 409 {object} model.ErrorResponse
// @Router /v1/register [post]
func RegisterController(c *fiber.Ctx) error {
	request := model.RegisterRequest{}
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	if err := controller.Register(request); err != nil {
		return err
	}

	return c.JSON(model.RegisterResponse{
		Ok:      true,
		Message: global.REGISTER_RESPONSE,
	})
}

// ValidateController godoc
// @Summary validate jwt token
// @Description validate jwt token
// @ID validate-jwt-token
// @Accept  json
// @Produce  json
// @Tags Validate
// @param x-jwt-assertion header string true "jwt header"
// @Success 200 {object} model.ValidateResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /v1/validate [get]
func ValidateController(c *fiber.Ctx) error {
	jwt := ""
	if jwt = c.Get(global.JWT_HEADER); jwt == "" {
		return error_lib.NoJwtInRequest
	}

	if err := controller.Validate(jwt); err != nil {
		return err
	}
	return c.JSON(model.ValidateResponse{
		Ok:      true,
		Message: global.VALIDATE_RESPONSE,
	})
}

// HealthController godoc
// @Summary health check
// @Description health check
// @ID health
// @Produce  json
// @Tags health
// @Success 200 {object} model.HealthResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /health [get]
func HealthController(c *fiber.Ctx) error {
	if err := controller.Health(); err != nil {
		return err
	}
	return c.JSON(model.HealthResponse{
		Ok:      true,
		Message: global.HEALTH_RESPONSE,
	})
}