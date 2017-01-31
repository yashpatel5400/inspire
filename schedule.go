package main

import (
	"fmt"
	"bufio"
	"os"
)

func regular() {
	return
}

func random() {
	return
}

func schedule() {
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
    	regular()

    case RANDOM:
    	random()

    case RETURN:
    	return
    }
}