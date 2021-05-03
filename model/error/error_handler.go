package error

import (
	"com.poalim.bank.hackathon.login-fiber/model"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func ErrorHandler(logger *zap.Logger) fiber.Config {
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
			reqInfo := model.RequestInfo{
				Method: c.Method(),
				Url:    string(c.Request().RequestURI()),
				Ip:     c.IP(),
				UID:    c.Get(fiber.HeaderXRequestID),
			}
			logger.Error(err.Error(), zap.Any("reqInfo", reqInfo))
			return c.Status(code).JSON(response)
		},
	}
}
