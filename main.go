package main

import (
	"log"
	"os"

	"github.com/brotherlogic/tasklister/server"
)

func writeString(filename, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		return err
	}
	return nil
}

func main() {
	log.Printf("Starting")
	s := server.NewServer()

	// Setup known_hosts
	err := writeString(os.Getenv("SSH_KNOWN_HOSTS"), os.Getenv("SSH_ENTRY"))
	if err != nil {
		log.Fatalf("Unable to write ssh entry: %v", err)
	}

	/* This piece is working:

	err = s.Test(os.Getenv("DEPLOY_KEY"))
	if err != nil {
		log.Fatalf("Error testing: % v -> %v", err, os.Getenv(("SSH_KNOWN_HOSTS")))
	}*/

	log.Printf("Tasklister started")

	// Keep the server running
	for true {
	}
}
