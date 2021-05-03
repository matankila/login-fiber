package service

import (
	"com.poalim.bank.hackathon.login-fiber/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func NewUuidMid() fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := uuid.New()
		c.Request().Header.Set(fiber.HeaderXRequestID, u.String())
		return c.Next()
	}
}

func NewLoggingMid(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestInfo{
			Method: c.Method(),
			Url:    string(c.Request().RequestURI()),
			Ip:     c.IP(),
			UID:    c.Get(fiber.HeaderXRequestID),
		}
		logger.Info("start",
			zap.Any("requestInfo", req),
			zap.String("uid", c.Get(fiber.HeaderXRequestID)))
		defer logger.Info("finish",
			zap.Any("requestInfo", req),
			zap.String("uid", c.Get(fiber.HeaderXRequestID)))
		return c.Next()
	}
}
