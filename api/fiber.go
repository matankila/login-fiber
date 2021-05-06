package api

import (
	"com.poalim.bank.hackathon.login-fiber/model"
	"com.poalim.bank.hackathon.login-fiber/service"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func InitApi(f *fiber.App, c Controller) {

	// middleware
	f.Use(service.NewUuidMid())
	f.Use(service.NewLoggingMid(service.GetLogger(service.Default)))

	// endpoints
	f.Get("/swagger/*", swagger.Handler)

	// endpoints grouped under /api
	api := f.Group("/api")
	api.Get("/health", c.health)
	// endpoints grouped under /api/v1
	v1 := api.Group("/v1")
	v1.Post("/login", c.login)
	v1.Post("/register", c.register)
	v1.Get("/validate", c.validate)
}

func ErrorHandler() fiber.Config {
	return fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			logger := service.GetLogger(service.Default)
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
