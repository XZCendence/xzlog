package xzlog

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	// Define log categories
	logger := &Logger{Output: os.Stdout, LogLevel: Verbose, LogToFile: true}
	Test := DeclareLogCategory("Test")
	logger.Log(Test, Verbose, "This is a verbose message")
	logger.Log(Test, Info, "This is an info message")
	logger.Log(Test, Debug, "This is a debug message")
	logger.Log(Test, Warning, "This is a warning message")
	logger.Log(Test, Error, "This is an error message")
	// This will exit, so be sure you don't use it somewhere you wouldn't use panic() or os.Exit():
	logger.Log(Test, Fatal, "This is a fatal message")
}
