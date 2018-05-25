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
	SleepUntil(time.Second * time.Duration(3))
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

// run using "go test -v"
