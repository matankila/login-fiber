package model

type AccountData struct {
	Id       string `json:"id" bson:"_id"`
	Password string `json:"password" bson:"password"`
}
