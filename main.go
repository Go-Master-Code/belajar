package logger

import "log"

func LogInfo(message string) {
	log.Printf("INFO - %v", message)
}
