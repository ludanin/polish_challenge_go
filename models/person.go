package models

import (
	"math/rand"
	"polish/utils"
	"strings"
	"time"
)

// Person represents a pointer on the map, it holds its own Coordinates
// and has functions to update its own position
type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	Coordinates Coordinates `json:"coordinates"`
}

// MakePerson generates a new Person with a random Name + random ID
func MakePerson() Person {
	// Generate ID
	id := strings.Builder{}
	idLength := 5

	r := rand.New(
		rand.NewSource(time.Now().UnixNano()),
	)

	charset := strings.Join(
		[]string{
			"abcdefghijklmnopqrstuvwxyz",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"1234567890",
		},
		"",
	)

	for i := 0; i < idLength; i++ {
		id.WriteByte(charset[r.Intn(len(charset))])
	}

	// Generate Name
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(names))

	// Offset from initial location
	coordinates := Coordinates{
		Latitude:  initialLocation.Latitude + (utils.RandomNoise() * 15),
		Longitude: initialLocation.Longitude + (utils.RandomNoise() * 15),
	}

	// Return the person
	return Person{
		ID:          id.String(),
		Name:        names[index],
		Coordinates: coordinates,
	}
}

var names = []string{
	"Dawn",
	"Carmon",
	"Vance",
	"Madeleine",
	"Erich",
	"Hassan",
	"Kenyetta",
	"Sherley",
	"Raven",
	"Adeline",
	"Carissa",
	"Jone",
	"Charlene",
	"Nita",
	"Belle",
	"Elsy",
	"Ofelia",
	"Claudine",
	"Tijuana",
	"Brynn",
	"Meaghan",
	"Anita",
	"Ira",
	"Ana",
	"Rosalina",
	"Leatrice",
	"Lonna",
	"Kacey",
	"Marinda",
	"Sharen",
	"Michale",
	"Clotilde",
	"Shaquita",
	"Libby",
	"Magaly",
	"Joanne",
	"Bettye",
	"Margarito",
	"Caryn",
	"Jacob",
	"Johnetta",
	"Sara",
	"Vanetta",
	"Roxann",
	"Hiedi",
	"Janna",
	"Clora",
	"Chun",
	"Yolonda",
	"Jacklyn",
}
