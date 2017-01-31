package main

import (
	"os"
	"time"
	"strconv"
)

/*****************************************************************************/
/* Produces a quick inspirational quote on time based on CMD-line args (secs)*/
/*****************************************************************************/
func main() {
	strint := os.Args[1]
	interval, err := strconv.Atoi(strint)
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Duration(interval) * time.Second)
		quick(false, false)
	}
}