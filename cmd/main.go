package main

import (
	"com.poalim.bank.hackathon.login-fiber/api"
	error_lib "com.poalim.bank.hackathon.login-fiber/error"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(error_lib.ErrorHandler())
	c := api.InitController()
	api.InitApi(app, c)
	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
