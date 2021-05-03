package api

import (
	"com.poalim.bank.hackathon.login-fiber/service"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func InitApi(f *fiber.App, c Controller) {

	f.Get("/swagger/*", swagger.Handler)
	swagger.New(swagger.Config{
		DeepLinking: false,
		URL:         "",
	})
	f.Use(service.NewUuidMid())
	f.Use(service.NewLoggingMid(service.GetLogger(service.Default)))
	api := f.Group("/api")
	api.Get("/health", c.health)
	v1 := api.Group("/v1")

	v1.Post("/login", c.login)
	v1.Post("/register", c.register)
	v1.Get("/validate", c.validate)
}
