package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Service   string `json:"service"`
	Message   string `json:"message"`
}

func writeLog(level, message string) {
	file, err := os.OpenFile("/var/log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Service:   "go-logging-service",
		Message:   message,
	}

	jsonLog, _ := json.Marshal(entry)

	file.Write(append(jsonLog, '\n'))
}

func main() {

	log.Println("Starting Go logging service...")

	for {
		writeLog("INFO", "User logged in")
		writeLog("ERROR", "Database connection failed")
		writeLog("DEBUG", "Processing request")

		time.Sleep(5 * time.Second)
	}
}
