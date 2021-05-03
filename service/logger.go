package service

import (
	"com.poalim.bank.hackathon.login-fiber/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger(logger Logger) *zap.Logger {
	return l[logger]
}

type loggerFactory map[Logger]*zap.Logger

type Logger interface {
	String() string
}

type DefaultLogger struct{}
type HealthLogger struct{}

var (
	l       = loggerFactory{}
	Default = DefaultLogger{}
	Health  = HealthLogger{}
	done    = make(chan struct{})
)

func (h HealthLogger) String() string {
	return "health"
}

func (d DefaultLogger) String() string {
	return "default"
}

func initLogger(loggerName string) *zap.Logger {
	lvl := zap.NewAtomicLevel()
	if err := lvl.UnmarshalText([]byte("info")); err != nil {
		panic(err)
	}
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	c := zap.NewProductionConfig()
	c.Level = lvl
	c.InitialFields = map[string]interface{}{"loggerName": loggerName}
	c.OutputPaths = []string{"stdout"}
	c.EncoderConfig = ec
	logger, err := c.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

// it inits the logger factory
// this function returns a channel, you must close before the program finishes.
func InitFactory() chan struct{} {
	l[Default] = initLogger(global.LOGGER_NAME)
	l[Health] = initLogger(global.HEALTH_LOGGER_NAME)
	// waits for channel to be closed and sync all loggers
	go func() {
		<-done
		for _, v := range l {
			v.Sync()
		}
	}()
	return done
}
