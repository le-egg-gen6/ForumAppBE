package utils

import (
	"forum/constant"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	UserID    uint
	Validated bool
}

func GenerateToken(userID uint, remember bool, validated bool) (string, error) {
	expiredTime := time.Now()
	if remember {
		expiredTime = expiredTime.Add(time.Hour * constant.ExpiredTimeInHourRemember)
	} else {
		expiredTime = expiredTime.Add(time.Hour * constant.ExpiredTimeInHour)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID:    userID,
		Validated: validated,
	})

	return token.SignedString([]byte(constant.SecretToken))
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(constant.SecretToken), nil
	})
}
