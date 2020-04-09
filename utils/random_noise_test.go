package utils_test

import (
	"polish/utils"
	"testing"
)

func TestRandomNoise(t *testing.T) {
	first := utils.RandomNoise()
	second := utils.RandomNoise()
	third := utils.RandomNoise()

	if first == second && third == second {
		t.Error("Could not generate random noise")
	}

	if first == 0 || second == 0 || third == 0 {
		t.Error("Random noise was zero")
	}
}
