package main

import (
	"com.poalim.bank.hackathon.login-fiber/api"
	"com.poalim.bank.hackathon.login-fiber/controller"
	"com.poalim.bank.hackathon.login-fiber/dao"
	"com.poalim.bank.hackathon.login-fiber/global/config"
	"com.poalim.bank.hackathon.login-fiber/service"
	"github.com/gofiber/fiber/v2"
	"time"

	_ "com.poalim.bank.hackathon.login-fiber/docs"
)

var (
	done  chan struct{}
	done2 chan struct{}
	db    dao.DB
)

func init() {
	config.InitConfig()
	done2 = service.InitLoggerFactory()
	if !config.Local {
		db, done = dao.New(config.ConnectionString)
	} else {
		db, done = dao.NewLocal()
	}
}

// @title Login
// @version 1.0
// @description Swagger for Login service
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email matan.k1500@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host
// @BasePath /api
func main() {
	app := fiber.New(api.ErrorHandler(service.GetLogger(service.Default)))
	h := service.NewHash()
	ctrlr := controller.NewController(db, h)
	c := api.NewApiController(ctrlr)
	api.InitApi(app, c)
	if err := app.Listen(":" + config.Port); err != nil {
		if !config.Local {
			close(done)
		}
		close(done2)
		time.Sleep(3 * time.Second)
		panic(err)
	}
}
