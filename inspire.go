package main

import (
	"bufio"
	"os"
	// "os/exec"
	"fmt"
)

func main() {
	QUICK	  := "1\n"
	SCHEDULE  := "2\n"
	CUSTOMIZE := "3\n"

	for {
		fmt.Println(`Please select from the following:
	(1) Quick
	(2) Schedule
	(3) Customize`)

		/* cmd := exec.Command("ls")
		err := cmd.Start()
		if err != nil {
			fmt.Println("Failed to start!")
			return
		} */

		reader  := bufio.NewReader(os.Stdin)
		mode, _ := reader.ReadString('\n')

		switch mode {
		case QUICK:
			if !quick(false) { return }

		case SCHEDULE:
			schedule()
			return

		case CUSTOMIZE:
			quick(true)
		}
	}
}