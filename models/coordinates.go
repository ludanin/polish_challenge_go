package models

import "polish/utils"

var initialLocation = Coordinates{
	Latitude: 59.955413, Longitude: 30.337844,
}

// Coordinates holds Latitude and Longitude values
type Coordinates struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}

// Noise adds some noise (i.e. very small decimal values, like 1/10000),
// to Latitude and Longitude
func (c *Coordinates) Noise() {
	(*c).Latitude += utils.RandomNoise()
	(*c).Longitude += utils.RandomNoise()
}
