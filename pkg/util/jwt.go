package util

import (
	"ginProject/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (token string, e error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginProject",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, e = tokenClaims.SignedString(jwtSecret)
	return
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, err := tokenClaims.Claims.(*Claims); err && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
