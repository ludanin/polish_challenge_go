package data

import (
	"polish/models"
	"time"
)

var people = models.MakePeople()

// GetPeople returns the current people on memory
func GetPeople() models.People { return people }

// ContinuousNoise ensures all the People inside our application keep moving
func ContinuousNoise() {
	for range time.Tick(1 * time.Second) {
		people.GenerateNoise()
	}
}
