package service

import (
	"log"
	"os"
)

type ServiceLogger struct {
	infoLogger *log.Logger
	errLogger  *log.Logger
}

const (
	ERROR  = 1
	INFO = 2
)

func (sl *ServiceLogger) PrintLog(message string, flag int) {
	if flag == ERROR {
		sl.errLogger.Println(message)
	} else if flag == INFO {
		sl.infoLogger.Println(message)
	}
}

func NewServiceLogger(infoLogFile *os.File, errLogFile *os.File) *ServiceLogger {
	infoLogger := log.New(infoLogFile, "INFO\t", log.Ldate|log.Ltime)
	errLogger := log.New(errLogFile, "ERROR\t", log.Ldate|log.Ltime)
	return &ServiceLogger{
		infoLogger: infoLogger,
		errLogger:  errLogger,
	}
}