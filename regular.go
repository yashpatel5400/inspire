package main

import (
	"os"
	"time"
	"strconv"
)

func main() {
	strint := os.Args[1]
	interval, err := strconv.Atoi(strint)
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Duration(interval) * time.Second)
		quick(false)
	}
}