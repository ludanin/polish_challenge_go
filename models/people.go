package models

import (
	"math/rand"
	"time"
)

// People is a group of Person
type People []Person

// MakePeople returns a slice full of Person, it will never be empty
func MakePeople() (result People) {
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

// GenerateNoise moves all the Person inside this People
func (p People) GenerateNoise() {
	for i := range p {
		p[i].Coordinates.Noise()
	}
}
