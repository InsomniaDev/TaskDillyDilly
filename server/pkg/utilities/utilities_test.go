package utilities

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestSleepUntil(t *testing.T) {
	t.Log("Starting a three second timer")
	startTime := time.Now()
	SleepUntil(Sleeper{time.Duration(3 * time.Second)})
	finishedTime := time.Now()

	if (finishedTime.Second() - startTime.Second()) == 3 {
		t.Log("Successfully waited for three seconds")
	} else {
		t.Errorf("Timer did not wait the appropriate time")
	}
}

func TestAssignLogFile(t *testing.T) {

	// Arrange
	logFile := "testing.log"
	textToFile := "Printed to file"

	// Act
	log.Println("Printed to console")
	err := AssignLogFile(logFile)
	if err != nil {
		t.Error(err)
	}
	log.Println(textToFile)

	// Assert
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Error("File does not exist")
	} else {
		t.Log("File exists")
		os.Remove(logFile)
		t.Log("Deleted file")
	}
}

func TestWorkerFile(t *testing.T) {
	firstTime := time.Now().Add(time.Second * 1)
	secondTime := time.Now().Add(time.Second * 3)

	first := Worker{"firstWorker", ":", SetSleeper(firstTime)}
	second := Worker{"secondWorker", ":", SetSleeper(secondTime)}

	Enqueue(first, second)

	StartProcessing()
}

// run using "go test -v"
