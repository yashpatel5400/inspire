package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	QUICK     := "1"
	SCHEDULE  := "2"
	CUSTOMIZE := "3"
	FAVORITES := "4"
	ABOUT     := "5"

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