// +build integration

package integration_test

import (
	"bytes"
	"com.poalim.bank.hackathon.login-fiber/global"
	"com.poalim.bank.hackathon.login-fiber/model"
	"encoding/json"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

/*
	This integration test can run against either local env or cloud env.
	for the test we use the test user (we are not creating new users but using this existing user)
	to run it you can type in your shell `make integration HOST=**`
*/

var host = flag.String("host", "http://localhost:8080", "host url")

func TestHealth(t *testing.T) {
	res, err := http.Get(*host + "/api/health")
	assert.Nil(t, err)
	defer res.Body.Close()
	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func TestLogin(t *testing.T) {
	r := model.LoginRequest{
		BankNumber: "test",
		AccountId:  "test",
		Password:   "test",
	}

	b, _ := json.Marshal(r)
	res, err := http.Post(*host+"/api/v1/login", fiber.MIMEApplicationJSON, bytes.NewBuffer(b))
	assert.Nil(t, err)
	defer res.Body.Close()
	assert.Equal(t, res.StatusCode, http.StatusOK)
	body, _ := ioutil.ReadAll(res.Body)
	login := model.LoginResponse{}
	json.Unmarshal(body, &login)
	assert.NotEmpty(t, login.Jwt)
	assert.True(t, login.Ok)
}

func TestRegister(t *testing.T) {
	r := model.RegisterRequest{
		BankNumber: "test",
		AccountId:  "test",
		Password:   "test",
	}

	b, _ := json.Marshal(r)
	res, err := http.Post(*host+"/api/v1/register", fiber.MIMEApplicationJSON, bytes.NewBuffer(b))
	assert.Nil(t, err)
	defer res.Body.Close()
	assert.Equal(t, res.StatusCode, http.StatusConflict)
}

func TestValidate(t *testing.T) {
	r := model.LoginRequest{
		BankNumber: "test",
		AccountId:  "test",
		Password:   "test",
	}

	b, _ := json.Marshal(r)
	res, err := http.Post(*host+"/api/v1/login", fiber.MIMEApplicationJSON, bytes.NewBuffer(b))
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	login := model.LoginResponse{}
	json.Unmarshal(body, &login)
	req, _ := http.NewRequest("GET", *host+"/api/v1/validate", nil)
	req.Header.Set(global.JWT_HEADER, login.Jwt)
	c := http.Client{}
	res2, err := c.Do(req)
	assert.Nil(t, err)
	defer res2.Body.Close()
	assert.Equal(t, res2.StatusCode, http.StatusOK)
}
