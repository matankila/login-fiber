package error

import (
	"com.poalim.bank.hackathon.login-fiber/model"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.Config {
	return fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			response := model.ErrorResponse{
				Ok:      false,
				Message: err.Error(),
			}
			return c.Status(code).JSON(response)
		},
	}
}
