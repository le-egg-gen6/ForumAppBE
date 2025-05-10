package utils

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
	"time"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GenerateValidateCode(numberDigits int) uint {
	if numberDigits < 1 || numberDigits > 18 { // Limit to prevent overflow
		return 0
	}

	rand.Seed(uint64(time.Now().UnixNano())) // Properly seed the random generator

	minimum := uint(1)
	for i := 1; i < numberDigits; i++ {
		minimum *= 10
	}
	maximum := minimum*10 - 1

	return minimum + uint(rand.Intn(int(maximum-minimum+1)))
}
