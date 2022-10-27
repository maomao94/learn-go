package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	findIds()
}

func findIds() {
	// find the process id by the process name
	fpid, err := robotgo.FindIds("Google")
	if err == nil {
		fmt.Println("pids...", fpid)
		if len(fpid) > 0 {
			robotgo.ActivePID(fpid[0])

			tl := robotgo.GetTitle(fpid[0])
			fmt.Println("pid[0] title is: ", tl)
		}
	}
}
