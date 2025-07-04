package logger

import (
	"log"
	"os"
)

var (
	logFile *os.File
)

func Init(logPath string) error {
	var err error
	logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(logFile)
	return nil
}

func Close() {
	if logFile != nil {
		logFile.Close()
	}
}

func Error(v ...interface{}) {
	log.SetPrefix("[ERROR] ")
	log.Println(v...)
}

func Info(v ...interface{}) {
	log.SetPrefix("[INFO] ")
	log.Println(v...)
}
