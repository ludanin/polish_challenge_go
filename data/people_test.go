package data_test

import (
	"polish/data"
	"sync"
	"testing"
	"time"
)

// Should move a person for about 6.5 seconds, if the person's coordinates
// don't update it's considered an error
func TestContinuousNoise(t *testing.T) {
	ticks := 0
	previousCoordinates := data.GetPeople()[0].Coordinates
	var wg sync.WaitGroup

	wg.Add(1)

	go data.ContinuousNoise()
	go func() {
		for range time.Tick(1050 * time.Millisecond) {
			if ticks >= 6 {
				wg.Done()
			}

			if ticks != 0 {
				if previousCoordinates == data.GetPeople()[0].Coordinates {
					t.Error("The Person didn't updated at ContinuousNoise")
				} else {
					previousCoordinates = data.GetPeople()[0].Coordinates
				}
			}

			t.Logf("%+v\n", data.GetPeople()[0])
			ticks++
		}
	}()

	wg.Wait()
}
