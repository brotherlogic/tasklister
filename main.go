package main

import (
	"log"
	"os"

	"github.com/brotherlogic/tasklister/server"
)

func main() {
	log.Printf("Starting")
	s := server.NewServer()
	err := s.Test(os.Getenv("DEPLOY_KEY"))
	if err != nil {
		log.Fatalf("Error testing: %v", err)
	}

	log.Printf("Server started successfully")

	// Keep the server running
	for true {
	}
}
