package main

import (
	"log"
	"strconv"

	"./pkg/utilities"
)

func main() {

	// Testing out functionality...

	// This changes all output to go to a log file rather than a console
	err := utilities.AssignLogFile("logfile.log")
	if err != nil {
		panic(err)
	}

	// Tests out the timer functionality
	for x := 0; x < 10; x++ {
		utilities.SecondSleeper()
		log.Println(strconv.Itoa(x+1) + " Timer has slept for a second")
	}
}
