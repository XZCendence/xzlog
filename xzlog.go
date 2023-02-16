package xzlog

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogLevel int

const (
	Verbose LogLevel = iota
	Debug
	Info
	Warning
	Error
	Fatal
)

var logLevelNames = map[LogLevel]string{
	Verbose: "Verbose",
	Debug:   "Debug",
	Info:    "Info",
	Warning: "Warning",
	Error:   "Error",
	Fatal:   "Fatal",
}

type Logger struct {
	Output      *os.File
	LogLevel    LogLevel
	LogToFile   bool
	LogFilePath string
}

var DefaultLogger = &Logger{
	Output:    os.Stderr,
	LogLevel:  Info,
	LogToFile: true, // Default to not log to a file
}

type LogCategory struct {
	Name string
}

func DeclareLogCategory(name string) *LogCategory {
	return &LogCategory{
		Name: name,
	}
}

// ANSI escape codes for colors
const (
	Cyan   = "\033[36m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
	Purple = "\033[35m"
)

func (l *Logger) Log(category *LogCategory, level LogLevel, v ...interface{}) {
	if level >= l.LogLevel {
		levelName := logLevelNames[level]
		now := time.Now().Format("2006-01-02 15:04:05.999")
		rawMessage := fmt.Sprintf("[%s] %s: %s: %v\n", now, category.Name, levelName, fmt.Sprint(v...))
		var cmessage string
		switch level {
		case Verbose:
			cmessage = Cyan + rawMessage + Reset
		case Debug:
			cmessage = Purple + rawMessage + Reset
		case Warning:
			cmessage = Yellow + rawMessage + Reset
		case Error:
			cmessage = Red + rawMessage + Reset
		case Fatal:
			cmessage = Red + rawMessage + "Exiting... " + Reset
			l.Output.Write([]byte(cmessage))
			os.Exit(1)
		}
		l.Output.Write([]byte(cmessage))
		// Wow that was some garbage tier golang, but it works
		// Now on to the logfile functionality

		// If the user specified their own path, create a file there
		// If not, create a file in the current directory
		if l.LogToFile {
			if l.LogFilePath == "" {
				// Default to the date
				l.LogFilePath = time.Now().Format("2006-01-02") + ".log"
			}
			f, err := os.OpenFile(l.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			if _, err := f.Write([]byte(rawMessage)); err != nil {
				log.Fatal(err)
			}
		}
	}
}

// Log with the default logger
func Log(level LogLevel, category *LogCategory, v ...interface{}) {
	DefaultLogger.Log(category, level, v...)
}

// There is no reason for this to exist, but I'm keeping it here for completeness
func SetDefaultLogLevel(level LogLevel) {
	DefaultLogger.LogLevel = level
}
