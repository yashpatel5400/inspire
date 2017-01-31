package main

import (
	"fmt"
	"bufio"
	"os"
	//"time"
)

func regular(interval float32) {
	return
}

func random() {
	return
}

func main() {
	REGULAR := "1\n"
	RANDOM  := "2\n"
	RETURN  := "3\n"
	
	fmt.Println(`Please select from the following:
(1) Regular
(2) Random
(3) Return`)
	reader  := bufio.NewReader(os.Stdin)
	mode, _ := reader.ReadString('\n')

	switch mode {
    case REGULAR:
    	regular(10.0)

    case RANDOM:
    	random()

    case RETURN:
    	return
    }
}