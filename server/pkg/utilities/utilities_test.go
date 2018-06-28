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

	expected := "Need to pass in value"
	actual := Enqueue()
	if actual.Error() != expected {
		t.Errorf("Error actual = %v, and Expected = %v.", actual, expected)
	}

	firstTime := time.Now().Add(time.Second * 1)
	secondTime := time.Now().Add(time.Second * 3)

	first := Worker{"firstWorker", "echo firstWorker", SetSleeper(firstTime)}
	second := Worker{"secondWorker", "echo secondWorker", SetSleeper(secondTime)}

	Enqueue(first, second)

	third := Worker{"thirdWorker", "echo thirdWorker", SetSleeper(firstTime)}
	fourth := Worker{"fourthWorker", "echo fourthWorker", SetSleeper(secondTime)}

	Enqueue(third, fourth)

	fifth := Worker{"fifthWorker", "open {Enter File Here}", SetSleeper(secondTime)}
	Enqueue(fifth)

	errs := StartProcessing()
	if errs != nil {
		errorFromProcessing := <- errs
		expected = "Test fifthWorker failed during execution"
		if errorFromProcessing.Error() != expected {
			t.Errorf("Error actual = %v, and Expected = %v.", actual, expected)
		}
	}
}

// run using "go test -v"
// For test coverage use "go test -coverprofile fmt"