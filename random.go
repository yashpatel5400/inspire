package main

import (
	"time"
	"math/rand"
)

func main() {
	// arbitrary interval length set as maximum
	MAX_INTERVAL := 4000
	rand.Seed(time.Now().UnixNano())
	interval := time.Duration(rand.Intn(MAX_INTERVAL))

	for {
		time.Sleep(interval * time.Second)
		quick(false)
		interval = time.Duration(rand.Intn(MAX_INTERVAL))
	}
}