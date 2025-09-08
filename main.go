package main

import (
	"log"
	"os"

	"example.com/go-polls/cmd/server"
)

func main() {
	// Check if .env file exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Println("Warning: .env file not found. Make sure to set REDIS_URL environment variable.")
	}

	// Start the server
	log.Println("Starting Go Polls server...")
	server.Run()
}
