package service

import (
	"log"
	"os"
)

type ServiceLogger struct {
	errLogger  *log.Logger
}

const (
	ERROR  = 1
)

func (sl *ServiceLogger) PrintLog(message string, flag int) {
	if flag == ERROR {
		sl.errLogger.Println(message)
	}
}

func NewServiceLogger(errLogFile *os.File) *ServiceLogger {
	errLogger := log.New(errLogFile, "ERROR\t", log.Ldate|log.Ltime)
	return &ServiceLogger{
		errLogger:  errLogger,
	}
}