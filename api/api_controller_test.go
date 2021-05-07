package api

import (
	"com.poalim.bank.hackathon.login-fiber/global"
	error_lib "com.poalim.bank.hackathon.login-fiber/global/error"
	"com.poalim.bank.hackathon.login-fiber/model"
	"com.poalim.bank.hackathon.login-fiber/service"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInitController(t *testing.T) {
	c := InitController()
	assert.NotNil(t, c)
}

func TestValidateController(t *testing.T) {
	app := fiber.New()
	app.Get("/validate", ValidateController)
	r := httptest.NewRequest("GET", "/validate", nil)
	r.Header.Set(global.JWT_HEADER, "x")
	res, err := app.Test(r)
	assert.Nil(t, err)
	assert.NotEqual(t, res.StatusCode, http.StatusOK)
}

func TestValidateController2(t *testing.T) {
	service.InitFactory()
	app := fiber.New(ErrorHandler(service.GetLogger(service.Default)))
	app.Get("/validate", ValidateController)
	res, err := app.Test(httptest.NewRequest("GET", "/validate", nil))
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
	r := model.ValidateResponse{
		Ok:      false,
		Message: error_lib.NoJwtInRequest.Message,
	}
	s, _ := json.Marshal(r)
	b, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, string(s), string(b))
}
