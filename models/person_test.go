package models_test

import (
	"polish/models"
	"testing"
)

// Should return 3 different persons
func TestMakePerson(t *testing.T) {
	first := models.MakePerson()
	second := models.MakePerson()
	third := models.MakePerson()

	if first.ID == second.ID || first.ID == third.ID {
		t.Error("Generating people with the same ID")
	}
	if first.Coordinates == second.Coordinates || first.Coordinates == third.Coordinates {
		t.Error("Generating people in the exact same location")
	}
}
