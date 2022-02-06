package logger

import "log"

func Info(str string) {
	log.Println("[Info]" + str)
}

func Error(str string) {
	log.Println("[Error]" + str)
}
