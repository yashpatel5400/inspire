package main

import (
	"bufio"
	"os"
	"os/exec"
	"syscall"
	"fmt"
	"strings"
)

/*****************************************************************************/
/* Given the name of an executable and the corresponding command line args,  */
/* forks into a parent and child process, runs sthe executable in background */
/* and returns the parent process (typically child will run indefinitely)    */
/*****************************************************************************/
func fork(filename string, args []string) {
	pid,_,execErr := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if execErr != 0 { panic(execErr) }
    switch (pid) {
        // child fork
        case 0:
        	binary, lookErr := exec.LookPath(filename)
			if lookErr != nil {
				panic(lookErr)
			}
			// fork new process and execute our program
			execErr := syscall.Exec(binary, args, os.Environ())

			// catch error if any
			if execErr != nil {
				panic(execErr)
			}
            return

        // parent fork
        default:
            return
    }
}

/*****************************************************************************/
/* Gives option to choose an inspirational quote (either favorited or not),  */
/* or to do so in the background, with options of either random or regular   */
/* intervals and providing corresponding interval breaks for the regular case*/
/*****************************************************************************/
func main() {
	QUICK	   := "1"
	FAVORITES  := "2"
	REGULAR    := "3"
	RANDOM     := "4"

	// fork new process and execute our program
	// execErr := syscall.Exec(binary, args, os.Environ())
	for {
		fmt.Println(`Please select from the following:
	(1) Quick
	(2) Favorites
	(3) Regular
	(4) Random`)

		reader  := bufio.NewReader(os.Stdin)
		mode, _ := reader.ReadString('\n')
		fmode   := strings.TrimSpace(mode)

		switch fmode {
		case QUICK:
			if !quick(false, true) { return }

		case FAVORITES:
			if !quick(true, true) { return }			

		case REGULAR:
			fmt.Println("Please enter your interval: ")
			interval, _ := reader.ReadString('\n')
			finterval   := strings.TrimSpace(interval)
			fork("./regular", []string{"./regular", finterval})
			return

		case RANDOM:
			fork("./random",  []string{"./random"})
			return
		}
	}
}