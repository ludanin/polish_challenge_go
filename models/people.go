package models

import (
	"math/rand"
	"time"
)

type people []Person

// MakePeople returns a slice full of Person, it will never be empty
func MakePeople() (result people) {
	min := 5
	max := 15

	rand.Seed(time.Now().UnixNano())
	amount := rand.Intn(max)
	if amount < min {
		amount = min
	}

	for i := 0; i <= amount; i++ {
		result = append(result, MakePerson())
	}

	return result
}

func (p people) GenerateNoise() {
	for i := range p {
		p[i].Coordinates.Noise()
	}
}
