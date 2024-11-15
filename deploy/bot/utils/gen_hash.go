package utils

import (
	"crypto/rand"
	"fmt"
	"time"
)

func GenerateHashCashChallenges() string {
	version := "1"
	difficulty := "20"
	timestamp := time.Now().Format("20060102")
	resource := "localhost"

	// Генерация случайной строки
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	randomString := fmt.Sprintf("%x", randomBytes)

	// Собираем задачу
	challenge := fmt.Sprintf("%s:%s:%s:%s::%s", version, difficulty, timestamp, resource, randomString)
	return challenge
}
