package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func Hash(text string) (string, error) {
	hasher := sha256.New()

	_, err := hasher.Write([]byte(text))
	if err != nil {
		return "", err
	}

	hashInBytes := hasher.Sum(nil)

	hash := hex.EncodeToString(hashInBytes)

	return hash, nil
}

func Compare(text string, hash string) (bool, error) {
	hashInBytes, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}

	hasher := sha256.New()

	_, err = hasher.Write([]byte(text))
	if err != nil {
		return false, err
	}

	hashed := hasher.Sum(nil)

	return hmac.Equal(hashed, hashInBytes), nil
}
