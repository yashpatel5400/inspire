package main

import (
	"time"
	"math/rand"
)

/*****************************************************************************/
/* Produces a quick inspirational quote at random interval btw 1-4000 seconds*/
/*****************************************************************************/
func main() {
	// arbitrary interval length set as maximum
	MAX_INTERVAL := 4000
	rand.Seed(time.Now().UnixNano())
	interval := time.Duration(rand.Intn(MAX_INTERVAL))

	for {
		time.Sleep(interval * time.Second)
		quick(false, false)
		interval = time.Duration(rand.Intn(MAX_INTERVAL))
	}
}