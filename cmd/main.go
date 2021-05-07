package main

import (
	"com.poalim.bank.hackathon.login-fiber/api"
	"com.poalim.bank.hackathon.login-fiber/service"
	"github.com/gofiber/fiber/v2"

	_ "com.poalim.bank.hackathon.login-fiber/docs"
)

func init() {
	service.InitFactory()
}

// @title Login
// @version 1.0
// @description Swagger for Login service
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email matan.k1500@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host
// @BasePath /api
func main() {
	app := fiber.New(api.ErrorHandler(service.GetLogger(service.Default)))
	c := api.InitController()
	api.InitApi(app, c)
	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
