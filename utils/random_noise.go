package utils

import (
	"math/rand"
	"time"
)

// RandomNoise returns a very small decimal variation, e.g. numbers close
// to 1/10000
//
// The value returned is never zero, but they can be negative
func RandomNoise() float64 {
	rand.Seed(time.Now().UnixNano())
	min := 0.00005
	number := (rand.Float64() * 10)
	if number < min {
		number = min
	}
	divisor := float64(10000)

	if rand.Float64() >= 0.5 {
		return number / divisor
	}

	return number / divisor * -1
}
