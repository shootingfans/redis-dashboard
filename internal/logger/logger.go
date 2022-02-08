package logger

import "log"

// Info is a simple logger (will refactor)
func Info(str string) {
	log.Println("[Info]" + str)
}

// Error is a simple logger (will refactor)
func Error(str string) {
	log.Println("[Error]" + str)
}
