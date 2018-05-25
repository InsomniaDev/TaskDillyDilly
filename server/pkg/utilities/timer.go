package utilities

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup
var completed bool

// Sleep for an hour
func HourSleeper() {
	time.Sleep(time.Hour)
}

// Sleep for a minute
func MinuteSleeper() {
	time.Sleep(time.Minute)
}

// Sleep for a second
func SecondSleeper() {
	time.Sleep(time.Second)
}

// Sleep until the provided time and return true when completed
func SleepUntil(setTime time.Duration) {
	time.Sleep(setTime)
}

// Possibly move to another package
func duration(year int, month time.Month, day, hour, min, sec, nsec int) time.Duration {
	t := time.Now()
	proposedTime := time.Date(year, month, day, hour, min, sec, nsec, t.Location())

	// Add another twenty four hours if the time has already passed
	if t.After(proposedTime) {
		proposedTime = proposedTime.Add(24 * time.Hour)
	}

	d := proposedTime.Sub(t)
	return d
}

func timerCompletion() {
	// defer wg.Done()
	completed = true
	log.Println("Timer has completed")
}
