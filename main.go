package main

import (
	"polish/data"
	"polish/server"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go data.ContinuousNoise()
	go server.Run()

	wg.Wait()
}
