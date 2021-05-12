package service

import "golang.org/x/crypto/bcrypt"

type Bycrypt interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type Hash struct{}

func NewHash() Bycrypt {
	return Hash{}
}

func (h Hash) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func (h Hash) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
