package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"learn.gin/pkg/setting"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password`
	jwt.StandardClaims
}

func GenerateToken(username string, password string) (string, error) {
	expireTime := time.Now().Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "learn-jwt",
		},
	}

	newClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := newClaims.SignedString(setting.JwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.JwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
