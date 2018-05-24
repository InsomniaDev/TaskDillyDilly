package utilities

import (
	"os"
	"syscall"
)

// Assign all output to log file for application
func AssignLogFile() {
	errorFile, err := os.OpenFile("LogFile", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	syscall.Dup2(int(errorFile.Fd()), 1)
	syscall.Dup2(int(errorFile.Fd()), 2)
}