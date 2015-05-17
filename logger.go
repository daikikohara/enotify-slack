package main

import (
	"log"
	"os"
)

// Logger holds log.Logger and os.File
type Logger struct {
	logger *log.Logger
	file   *os.File
}

// NewLogger constructs Logger and returns it.
// This function opens a log file defined in the configuration file(yaml) and create log.Logger according to the log file.
func NewLogger(path string) *Logger {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panic("can't open", path)
	}
	logger := log.New(file, "event-notify-log: ", log.Lshortfile|log.LstdFlags)
	return &Logger{logger, file}
}

// Close closes a file held by Logger.
func (self *Logger) Close() {
	err := self.file.Close()
	if err != nil {
		log.Println(err.Error())
	}
}

// Println prints a message to console and Logger.file
func (self *Logger) Println(msg string) {
	log.Println(msg)
	self.logger.Println(msg)
}
