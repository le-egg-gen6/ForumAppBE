package util

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const SecretToken = "ledeptraivailzz"
const ExpiredTimeInHour = 24
const ExpiredTimeInHourRemember = 24 * 7

func GenerateToken(userID uint64, remember bool) (string, error) {
	expiredTime := time.Now()
	if remember {
		expiredTime.Add(time.Hour * ExpiredTimeInHourRemember)
	} else {
		expiredTime.Add(time.Hour * ExpiredTimeInHour)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": expiredTime.Unix(),
	})

	return token.SignedString([]byte(SecretToken))
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(SecretToken), nil
	})
}
