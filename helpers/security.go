package helpers

import (
	"crypto/sha512"
	"encoding/hex"
	"time"
)

func GenerateTokenRecovery() string {
	today := time.Now()
	k := today.String() + "Presto"
	h := sha512.New()
	h.Write([]byte(k))
	return hex.EncodeToString(h.Sum(nil))
}
