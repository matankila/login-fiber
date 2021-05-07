package service

import (
	error_lib "com.poalim.bank.hackathon.login-fiber/global/error"
	"com.poalim.bank.hackathon.login-fiber/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	singletonJwtWrapper = &jwtWrapper{
		SecretKey:       "ttlogin",
		Issuer:          "ttlogin",
		ExpirationHours: 72,
	}
)

// jwtWrapper wraps the signing key and the issuer
type jwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// JwtClaim adds email as a claim to the token
type JwtClaim struct {
	BankNumber string
	AccountId  string
	jwt.StandardClaims
}

func NewJwtWrapper() *jwtWrapper {
	return singletonJwtWrapper
}

// GenerateToken generates a jwt token
func (j *jwtWrapper) GenerateToken(request model.LoginRequest) (signedToken string, err error) {
	claims := &JwtClaim{
		BankNumber: request.BankNumber,
		AccountId:  request.AccountId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}
	return
}

//ValidateToken validates the jwt token
func (j *jwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, _ = token.Claims.(*JwtClaim)
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, error_lib.ExpiredJwt
	}
	return
}
