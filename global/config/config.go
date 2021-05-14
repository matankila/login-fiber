package config

import (
	"com.poalim.bank.hackathon.login-fiber/global"
	"flag"
	"os"
	"strconv"
)

var (
	Local            bool
	ConnectionString string
	Port             string
	LogLvl           string
)

func InitConfig() {
	flag.BoolVar(&Local, "local", getEnvBool("LOCAL", true), "is running on Local machine")
	flag.StringVar(&ConnectionString, "conn", getEnvStr("CONN", global.URI), "db connection string")
	flag.StringVar(&Port, "port", getEnvStr("PORT", "8080"), "application Port")
	flag.StringVar(&LogLvl, "logLvl", getEnvStr("LOG_LVL", "info"), "application log level")
	flag.Parse()
}

func getEnvStr(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if v, ok := os.LookupEnv(key); ok {
		if b, err := strconv.ParseBool(v); err == nil {
			return b
		}
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v, ok := os.LookupEnv(key); ok {
		if b, err := strconv.Atoi(v); err == nil {
			return b
		}
	}

	return fallback
}
