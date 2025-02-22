package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	debugMode bool
}

func New(debugMode bool) *Logger {
	return &Logger{debugMode: debugMode}
}

func (l *Logger) Info(format string, args ...interface{}) {
	timestamp := time.Now().Format("15:04:05")
	message := fmt.Sprintf(format, args...)
	fmt.Printf("[%s] [INFO] %s\n", timestamp, message)
}

func (l *Logger) Error(format string, args ...interface{}) {
	timestamp := time.Now().Format("15:04:05")
	message := fmt.Sprintf(format, args...)
	fmt.Printf("[%s] [ERROR] %s\n", timestamp, message)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	if !l.debugMode {
		return
	}

	timestamp := time.Now().Format("15:04:05")
	message := fmt.Sprintf(format, args...)
	fmt.Printf("[%s] [DEBUG] %s\n", timestamp, message)
}
