package api

import (
	"github.com/gofiber/fiber/v2"
)

//TODO: middleware + error handler

func InitApi(f *fiber.App, c Controller) {

	api := f.Group("/api")
	api.Get("/health", c.health)
	v1 := api.Group("/v1")

	v1.Post("/login", c.login)
	v1.Post("/register", c.register)
	v1.Get("/validate", c.validate)
}
