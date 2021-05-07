package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNewLoggingMid(t *testing.T) {
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
	app := fiber.New()
	app.Use(NewLoggingMid(l))
	app.Test(httptest.NewRequest("GET", "/", nil))
	d, err := ioutil.ReadFile("./logger.txt")
	assert.Nil(t, err)
	assert.NotEmpty(t, string(d))

}

func TestNewUuidMid(t *testing.T) {
	f := fiber.New()
	f.Use(NewUuidMid())
	resp, _ := f.Test(httptest.NewRequest("GET", "/", nil))
	uid := resp.Header.Get(fiber.HeaderXRequestID)
	assert.NotEmpty(t, uid)
}
