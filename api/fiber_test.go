package api

import (
	error_lib "com.poalim.bank.hackathon.login-fiber/global/error"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	f, err := os.Create("logger.txt")
	assert.Nil(t, err)
	defer f.Close()
	defer os.Remove("./logger.txt")
	f.Chmod(0755)
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./logger.txt",
	}
	l, err := cfg.Build()
	assert.Nil(t, err)
	app := fiber.New(ErrorHandler(l))
	app.Get("/", func(ctx *fiber.Ctx) error {
		return error_lib.NoJwtInRequest
	})
	app.Test(httptest.NewRequest("GET", "/", nil))
	d, err := ioutil.ReadFile("./logger.txt")
	assert.Nil(t, err)
	assert.NotEmpty(t, string(d))
}
