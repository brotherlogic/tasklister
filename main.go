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
		log.Fatalf("Error testing: % v -> %v", err, os.Getenv(("SSH_KNOWN_HOSTS")))
	}

	log.Printf("Tasklister started")

	// Keep the server running
	for true {
	}
}
