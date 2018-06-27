package main

import (
	"log"
	"strconv"

	utilities "github.com/InsomniaDev/TaskDillyDilly/server/pkg/utilities"
)

func main() {

	// Testing out functionality...
	// Below commented out code needs updated utility library in github
	// current := time.Now().Add(time.Second * 10)
	// job := utilities.Worker{"Test Job", current}
	// utilities.Enqueue(job)
	// utilities.StartProcessing()

	// Tests out the timer functionality
	for x := 0; x < 10; x++ {
		utilities.SecondSleeper()
		log.Println(strconv.Itoa(x+1) + " Timer has slept for a second")
	}
}
