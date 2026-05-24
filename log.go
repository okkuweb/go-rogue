package main

import (
	"log"
	"os"
)

var logger *log.Logger
var logFile *os.File

func InitLogger() {
	logFile, err := os.OpenFile("./app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	logger = log.New(logFile, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Log(args ...any) {
	if logger == nil {
		log.Println("Logger not initialized!")
		return
	}
	logger.Println(args...)
}

func DebugLog(args ...any) {
	if logger == nil {
		log.Println("Logger not initialized!")
		return
	}
	logger.Println(args...)
}

func GameLog(args ...any) {
	if logger == nil {
		log.Println("Logger not initialized!")
		return
	}
	args = append([]any{"GAME: "}, args...)
	logger.Println(args...)
	// TODO: Game message log print function here
}
