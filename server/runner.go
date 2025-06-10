package server

import (
	"log"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Test() error {
	// Get a tmp directory
	log.Printf("Checking out")
	tDir, err := os.MkdirTemp("", "gitcheckout")
	if err != nil {
		return err
	}

	t1 := time.Now()
	_, err = git.PlainClone(tDir, false, &git.CloneOptions{
		URL: "https://github.com/brotherlogic/tasklister",
	})
	log.Printf("Took %v to clone", time.Since(t1))

	return err
}
