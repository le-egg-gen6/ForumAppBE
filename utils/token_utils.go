package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
)

func TokenHash(token string) string {

	hasher := md5.New()
	hasher.Write([]byte(token))

	theHash := hex.EncodeToString(hasher.Sum(nil))

	u := uuid.New()

	theToken := theHash + u.String()
	return theToken

}
