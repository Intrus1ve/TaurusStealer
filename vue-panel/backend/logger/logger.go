package logger

import (
	"log"
	"os"
)

func Log() *log.Logger {
	logPath := "./panel_errors.log"
	var logFile *os.File
	_, err := os.Stat(logPath)
	if os.IsNotExist(err) {
		logFile, _ = os.Create(logPath)
	} else {
		logFile, _ = os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0600)
	}

	return log.New(logFile, "", log.LstdFlags)
}
