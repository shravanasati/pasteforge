package utils

import (
	"math/rand"
	"strings"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890_"

func GenerateRandomID(length int) string {
	if length <= 0 {
		panic("generate random id: invalid len passed <= 0")
	}

	var s strings.Builder
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	total := len(charset)

	for i := 0; i < length; i++ {
		s.WriteByte(charset[rng.Intn(total)])
	}

	return s.String()
}
