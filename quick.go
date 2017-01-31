package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"math/rand"
	"time"
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
	numLines := len(lines)
	
	if numLines == 0 {
		return false
	}

	rand.Seed(time.Now().UnixNano())
	myInspiration := lines[rand.Intn(numLines)]
	fmt.Println(myInspiration)
	return true
}