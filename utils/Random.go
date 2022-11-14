package utils

import (
	"math/rand"
	"time"
)

func RandomNumber(maxRange int) int {
	randomx := rand.NewSource(time.Now().UnixNano())
  random := rand.New(randomx)

	return random.Intn(maxRange)
}

// https://www.geeksforgeeks.org/generating-random-numbers-in-golang/