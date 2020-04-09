package models_test

import (
	"polish/models"
	"testing"
)

func TestMakePeople(t *testing.T) {
	people := models.MakePeople()

	if len(people) == 0 {
		t.Error("MakePeople should never generate an empty slice")
	}

	// Check if a duplicated ID was found
	ids := make(map[string]bool)
	for _, p := range people {
		if ids[p.ID] {
			t.Error("MakePeople should never generate Person with duplicated IDs")
		} else {
			ids[p.ID] = true
		}
	}
}

func TestPeopleGenerateNoise(t *testing.T) {
	people := models.MakePeople()

	initial := people[0].Coordinates

	t.Logf("%+v\n", people[0])
	people.GenerateNoise()

	if initial == people[0].Coordinates {
		t.Error("People.GenerateNoise should move everyone")
	}
}
