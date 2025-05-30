package utils

import (
	"forum/constant"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func GenerateToken(userID uint, remember bool) (string, error) {
	expiredTime := time.Now()
	if remember {
		expiredTime = expiredTime.Add(time.Hour * constant.ExpiredTimeInHourRemember)
	} else {
		expiredTime = expiredTime.Add(time.Hour * constant.ExpiredTimeInHour)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(int(userID)),
		ExpiresAt: jwt.NewNumericDate(expiredTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	return token.SignedString([]byte(constant.SecretToken))
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(constant.SecretToken), nil
	})
}
