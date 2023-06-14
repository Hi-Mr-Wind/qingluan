package jwtutil

import (
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"waveQServer/src/utils/logutil"
)

const (
	KEY     string = "qingluan message queue"
	OutTime        = time.Hour * 2 //默认过期时间
)

var secret = []byte(base64.StdEncoding.EncodeToString([]byte(KEY)))

// GetToken 获取令牌
func GetToken(username string, password string) (string, error) {

	c := jwt.StandardClaims{
		Id:        base64.StdEncoding.EncodeToString([]byte(username + password)),
		ExpiresAt: time.Now().Add(OutTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedString, err := token.SignedString(secret)
	return signedString, err
}

// ParseToken 解析令牌
func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		logutil.LogError(err.Error())
		return nil, err
	}
	if token.Valid {
		return nil, errors.New("login expired")
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if ok {
		return claims, nil
	}
	return nil, errors.New("the token is not valid")
}
