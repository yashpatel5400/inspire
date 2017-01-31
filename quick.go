package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"math/rand"
	"time"
	"color"
)

func readFile(filename string) ([]string) {
	content, err := ioutil.ReadFile(filename)
	var fileStrings []string
	if err != nil {
		fmt.Println("Failed to read file!")
		return fileStrings
	}
	return strings.Split(string(content), "\n")
}

func quick() (bool) {
	lines := readFile("list.txt")
	colors := [...]*color.Color{color.New(color.FgBlue), color.New(color.FgGreen),
		color.New(color.FgYellow), color.New(color.FgMagenta),
		color.New(color.FgCyan), color.New(color.FgWhite)}

	if len(lines) == 0 {
		return false
	}

	rand.Seed(time.Now().UnixNano())
	printColor    := colors[rand.Intn(len(colors))].Add(color.Bold)
	myInspiration := lines[rand.Intn(len(lines))]
	printColor.Println(myInspiration)
	return true
}