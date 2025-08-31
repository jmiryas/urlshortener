package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

const TOKEN_LENGTH = 8

func GenerateToken(url string) string {
	hash := sha256.Sum256([]byte(url))
	
	// Convert to hexadecimal string and take first 8 characters
	token := hex.EncodeToString(hash[:])
	return token[:TOKEN_LENGTH]
}