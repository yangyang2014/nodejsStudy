package logger

import (
	"log"
)

func Info(format string, v ...interface{}) {
	log.Printf(format+"\n", v...)
}
func Debug(format string, v ...interface{}) {
	log.Printf(format+"\n", v...)
}
func Warn(format string, v ...interface{}) {
	log.Printf(format+"\n", v...)
}
func Error(format string, v ...interface{}) {
	log.Printf(format+"\n", v...)
}
