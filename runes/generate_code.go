package runes

import (
	"fmt"
	"math/rand"
	"time"
)

var digitsAndNumbers = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomCode(codeLength int) string {
	code := make([]rune, codeLength)
	fmt.Println("code", code)
	for i := range code {
		code[i] = digitsAndNumbers[rand.Intn(len(digitsAndNumbers))]
		fmt.Println("code[", i, "]:", code[i])
	}
	return string(code)
}
