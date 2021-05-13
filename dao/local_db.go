package dao

import (
	error_lib "com.poalim.bank.hackathon.login-fiber/global/error"
	"com.poalim.bank.hackathon.login-fiber/model"
	"errors"
	"sync"
)

type localDb struct {
	m sync.Map
}

func NewLocal() (DB, chan struct{}) {
	return &localDb{
		m: sync.Map{},
	}, nil
}

func (l *localDb) Get(requestObj interface{}) (interface{}, error) {
	req, ok := requestObj.(model.AccountData)
	if !ok {
		return nil, error_lib.UnsupportedType
	}

	res, ok := l.m.Load(req.Id)
	if !ok {
		return nil, errors.New("not found")
	}

	return res.(model.AccountData).Password, nil
}

func (l *localDb) Set(requestObj interface{}) error {
	req, ok := requestObj.(model.AccountData)
	if !ok {
		return error_lib.UnsupportedType
	}

	l.m.Store(req.Id, req)

	return nil
}

func (l *localDb) Ping() (bool, error) {
	return true, nil
}
