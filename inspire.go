package main

import (
	"bufio"
	"os"
	"os/exec"
	"syscall"
	"fmt"
)

func main() {
	QUICK	  := "1\n"
	SCHEDULE  := "2\n"

	pid,_,execErr := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if execErr != 0 { panic(execErr) }
    switch (pid) {
        // child fork
        case 0:
        	binary, lookErr := exec.LookPath("./schedule")
			if lookErr != nil {
				panic(lookErr)
			}
			// fork new process and execute our program
			execErr := syscall.Exec(binary, []string{}, os.Environ())

			// catch error if any
			if execErr != nil {
				panic(execErr)
			}
            return

        // parent fork
        default:
            fmt.Println("PARENT")
            return
    }

	// fork new process and execute our program
	// execErr := syscall.Exec(binary, args, os.Environ())
	for {
		fmt.Println(`Please select from the following:
	(1) Quick
	(2) Schedule`)

		reader  := bufio.NewReader(os.Stdin)
		mode, _ := reader.ReadString('\n')

		switch mode {
		case QUICK:
			if !quick(false) { return }

		case SCHEDULE:
			return

		}
	}
}