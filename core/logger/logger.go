package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func initialise(infoHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle, "", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "", log.Ldate|log.Ltime|log.Lshortfile)

}

func init() {
	infoLog := createLogFile("logs", "info")
	errorLog := createLogFile("logs", "error")

	infoWriter := io.MultiWriter(infoLog, os.Stdout)
	errorWriter := io.MultiWriter(errorLog, os.Stderr)

	initialise(infoWriter, errorWriter)

}

func createLogFile(path string, logType string) *os.File {
	file, err := os.OpenFile(fmt.Sprintf("%s/%s_%s.log", path, logType, time.Now().UTC().Format("2006-01-02")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error creating log file : %s", err)
	}
	return file
}
