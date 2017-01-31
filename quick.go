package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"io/ioutil"
	"math/rand"
	"time"
	"color"
)
	
/*****************************************************************************/
/* Checks to see if there was an error and panics with it if present         */
/*****************************************************************************/
func check(e error) {
    if e != nil {
        panic(e)
    }
}

/*****************************************************************************/
/* Given a filename string, returns the array of strings contained in file.  */
/* If none present (i.e. error occurs), alerts user and returns empty array  */
/*****************************************************************************/
func readFile(filename string) ([]string) {
	content, err := ioutil.ReadFile(filename)
	var fileStrings []string
	if err != nil {
		fmt.Println("Failed to read file! Must not exist or have errors!")
		return fileStrings
	}
	return strings.Split(string(content), "\n")
}

/*****************************************************************************/
/* Produces a quote of inspiration. Returns a boolean if wanting to retry    */
/*****************************************************************************/
func quick(isFavorites bool) (bool) {
	var lines []string
	if (isFavorites) {
		lines = readFile("favorites.txt")
	} else {
		lines = readFile("list.txt")	
	}
	
	colors := [...]*color.Color{
		color.New(color.FgBlue), color.New(color.FgGreen),
		color.New(color.FgYellow), color.New(color.FgMagenta),
		color.New(color.FgCyan), color.New(color.FgWhite)}

	if len(lines) == 0 {
		return false
	}

	rand.Seed(time.Now().UnixNano())
	printColor    := colors[rand.Intn(len(colors))].Add(color.Bold)
	myInspiration := lines[rand.Intn(len(lines))]
	printColor.Println(myInspiration)
	
	FAVORITE := "1\n"
	RETRY    := "2\n"

	fmt.Println(`Please select from the following:
	(1) Add to Favorites
	(2) Retry
	(3) Quit`)
	reader  := bufio.NewReader(os.Stdin)
	mode, _ := reader.ReadString('\n')

	if mode == FAVORITE {
		var f *os.File
		var err error
		if _, err := os.Stat("./favorites.txt"); os.IsNotExist(err) {
			f, err = os.Create("./favorites.txt")
		} else {
			f, err = os.Open("./favorites.txt")
		}
		
		check(err)
		defer f.Close()

	    _, err = f.WriteString(myInspiration)
	    check(err)
	    f.Sync()
	    return true
	}

	return (mode == RETRY)
}