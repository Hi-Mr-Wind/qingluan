package jwtutil

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
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

func RenewToken(parsedToken *jwt.StandardClaims) (string, error) {
	// 检查 Token 是否已过期
	if parsedToken.ExpiresAt > time.Now().Unix() {
		// 进行续签操作
		// 更新响应的过期时间
		newExpiresAt := time.Now().Add(OutTime).Unix()
		parsedToken.ExpiresAt = newExpiresAt
		// 生成新的 Token
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, parsedToken)
		// 签名 Token
		tokenString, err := newToken.SignedString(secret)
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
	// 如果 Token 尚未过期，则无需续签
	return "", nil
}

func GenerateToken() string {
	// 生成UUID作为唯一标识符
	uuid := uuid.New().String()

	// 获取当前时间戳
	timestamp := time.Now().Unix()

	// 将UUID和时间戳连接起来
	data := fmt.Sprintf("%s-%d", uuid, timestamp)

	// 对数据进行哈希（使用MD5算法作为示例）
	hasher := md5.New()
	hasher.Write([]byte(data))
	hash := hex.EncodeToString(hasher.Sum(nil))

	// 返回哈希缩短的token
	return hash
}
