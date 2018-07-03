//+build unix, darwin

package utilities

import (
	"log"
	"os"
	"syscall"
)

var logFile string

// Assign all output to log file for application
func AssignLogFile(file string) error {
	errorFile, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	errorFile.Close()

	syscall.Dup2(int(errorFile.Fd()), 1)
	syscall.Dup2(int(errorFile.Fd()), 2)
	log.Println("Began logging to file")
	defer errorFile.Close()
	return nil
}

// Notes: This will save the log file wherever specified,
// 		Otherwise it will create a log file wherever the application is called from
