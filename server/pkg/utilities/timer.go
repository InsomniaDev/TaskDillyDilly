package utilities

import (
	"sync"
	"time"
)

type Sleeper struct {
	executionTime time.Duration
}

var wg sync.WaitGroup

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
func SleepUntil(s Sleeper) {
	time.Sleep(s.executionTime)
}

// Possibly move to another package
func SetSleeper(proposedTime time.Time) Sleeper {
	t := time.Now()

	// Add another twenty four hours if the time has already passed
	if t.After(proposedTime) {
		proposedTime = proposedTime.Add(24 * time.Hour)
	}

	d := proposedTime.Sub(t)
	return Sleeper{d}
}
