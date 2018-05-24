package main

import (
	"fmt"
	"strconv"

	"./pkg/utilities"
)

func main() {

	// Testing out functionality...

	// This changes all output to go to a log file rather than a console
	utilities.AssignLogFile()

	// Tests out the timer functionality
	for x := 0; x < 10; x++ {
		utilities.SecondSleeper()
		fmt.Println(strconv.Itoa(x+1) + " Timer has slept for a second")
	}
}
