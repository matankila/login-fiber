package dao

import (
	"com.poalim.bank.hackathon.login-fiber/model"
	error_lib "com.poalim.bank.hackathon.login-fiber/model/error"
)

var (
	db = map[string]string{}
)

func SetAccount(account model.AccountData) error {
	db[account.Id] = account.Password
	return nil
}

func GetAccount(account model.AccountData) (string, error) {
	if v, ok := db[account.Id]; !ok {
		return "", error_lib.AccountNotExists
	} else {
		return v, nil
	}
}
func Ping() (bool, error) {
	return true, nil
}