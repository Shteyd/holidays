package token

import (
	"context"
	"math/rand"
	"unsafe"
)

const (
	letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

type TokenManager struct {
	length uint16
}

func New(
	tokenLength uint16,
) TokenManager {
	return TokenManager{
		length: tokenLength,
	}
}

func (token TokenManager) GenerateToken(_ context.Context) (string, error) {
	tokenSlice := make([]byte, 0, token.length)
	for range tokenSlice {
		tokenSlice = append(tokenSlice, letterRunes[rand.Intn(len(letterRunes))])
	}

	return unsafe.String(unsafe.SliceData(tokenSlice), len(tokenSlice)), nil
}
