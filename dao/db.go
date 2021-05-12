package dao

import (
	error_lib "com.poalim.bank.hackathon.login-fiber/global/error"
	"com.poalim.bank.hackathon.login-fiber/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

var (
	m    = mongoDB{}
	done = make(chan struct{})
)

type mongoDB struct {
	DB *mongo.Database
	sync.Once
}

type DB interface {
	Set(interface{}) error
	Get(interface{}) (interface{}, error)
	Ping() (bool, error)
}

func New(uri string) (DB, chan struct{}) {
	m.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}

		go func() {
			<-done
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()

		m.DB = client.Database("login")
	})

	return &m, done
}

func (m *mongoDB) Get(requestObj interface{}) (interface{}, error) {
	req, ok := requestObj.(model.AccountData)
	if !ok {
		return nil, error_lib.UnsupportedType
	}

	collection := m.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	res := bson.M{}
	err := collection.FindOne(ctx, bson.M{"_id": req.Id}).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res["password"], nil
}

func (m *mongoDB) Set(requestObj interface{}) error {
	req, ok := requestObj.(model.AccountData)
	if !ok {
		return error_lib.UnsupportedType
	}

	collection := m.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoDB) Ping() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := m.DB.Client().Ping(ctx, readpref.Primary()); err != nil {
		return false, err
	}

	return true, nil
}
