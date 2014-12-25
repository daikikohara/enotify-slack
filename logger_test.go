package main_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	. "github.com/daikikohara/enotify-slack"
)

const logMessage = "test message"

func TestPrintln(t *testing.T) {
	logFile, err := ioutil.TempFile(os.TempDir(), "event-notify-test.log")
	if err != nil {
		t.Error(err.Error())
	}
	defer os.Remove(logFile.Name())

	logger := NewLogger(logFile.Name())
	logger.Println(logMessage)

	b, err := ioutil.ReadFile(logFile.Name())
	if err != nil {
		t.Error(err.Error())
	}
	if !strings.Contains(string(b), logMessage) {
		t.Error("log file doesn't contain log message")
	}

	logger.Close()
}
