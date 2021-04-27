package api

import (
	"com.poalim.bank.hackathon.login-fiber/controller"
	error_lib "com.poalim.bank.hackathon.login-fiber/error"
	"com.poalim.bank.hackathon.login-fiber/global"
	"com.poalim.bank.hackathon.login-fiber/model"
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
	c.login = loginController
	c.register = registerController
	c.validate = validateController
	c.health = healthController
	return c
}

func loginController(c *fiber.Ctx) error {
	request := model.LoginRequest{}
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	response, err := controller.Login(request)
	if err != nil {
		return err
	}
	jwt, err := controller.CreateJwt(request)
	if err != nil {
		return err
	}
	cookie := fiber.Cookie{
		Name:  "jwt",
		Value: jwt,
	}
	c.Cookie(&cookie)
	return c.JSON(response)
}

func registerController(c *fiber.Ctx) error {
	request := model.RegisterRequest{}
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	//TODO: business logic
	response, err := controller.Register(request)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func validateController(c *fiber.Ctx) error {
	jwt := ""
	if jwt = c.Get(global.JWT_HEADER); jwt == "" {
		return error_lib.NoJwtInRequest
	}

	response, err := controller.Validate(jwt)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func healthController(c *fiber.Ctx) error {
	response, err := controller.Health()
	if err != nil {
		return err
	}

	return c.JSON(response)
}
