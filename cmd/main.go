package main

import (
	"com.poalim.bank.hackathon.login-fiber/api"
	error_lib "com.poalim.bank.hackathon.login-fiber/model/error"
	"com.poalim.bank.hackathon.login-fiber/service"
	"github.com/gofiber/fiber/v2"

	_ "com.poalim.bank.hackathon.login-fiber/docs"
)

func init() {
	service.InitFactory()
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host
// @BasePath /api
func main() {
	logger := service.GetLogger(service.Default)
	app := fiber.New(error_lib.ErrorHandler(logger))
	c := api.InitController()
	api.InitApi(app, c)
	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
