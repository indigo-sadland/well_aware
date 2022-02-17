package utils

import (
	"log"
	"os"
	"runtime"
)

var ErrorLogger   *log.Logger

// ErrorLogInit creates logger for errors.
func ErrorLogInit() {

	var logFile string

	dir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	if runtime.GOOS == "windows" {
		logFile = dir + "\\" + "well_aware_logs.txt"
	} else {
		logFile = dir + "/" + "well_aware_logs.txt"
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}