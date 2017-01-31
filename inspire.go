package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	QUICK     := "1\n"
	SCHEDULE  := "2\n"
	CUSTOMIZE := "3\n"
	FAVORITES := "4\n"
	ABOUT     := "5\n"

	fmt.Println(`Please select from the following:
		(1) Quick
		(2) Schedule
		(3) Customize
		(4) Favorites
		(5) About`)

	reader  := bufio.NewReader(os.Stdin)
	mode, _ := reader.ReadString('\n')

	switch mode {
    case QUICK:
    	quick()

    case SCHEDULE:
    	quick()

	case CUSTOMIZE:
    	quick()

	case FAVORITES:
    	quick()

	case ABOUT:
   		quick()
    }
}