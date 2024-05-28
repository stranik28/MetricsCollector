package server

import (
	"crypto/hmac"
	"crypto/sha256"
)

func Signature(body []byte, key []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(body)
	return hash.Sum(nil)
}
