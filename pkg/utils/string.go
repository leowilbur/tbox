package utils

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// MakeRandString make random string
func MakeRandString(n int, arr []rune) string {
	if arr == nil {
		arr = letterRunes
	}
	b := make([]rune, n)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = arr[rand.Intn(len(arr))]
	}
	return string(b)
}
