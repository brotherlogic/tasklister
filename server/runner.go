package server

import (
	"os"

	"github.com/go-git/go-git/v5"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Test() error {
	// Get a tmp directory
	tDir, err := os.MkdirTemp("", "gitcheckout")
	if err != nil {
		return err
	}
	_, err = git.PlainClone(tDir, false, &git.CloneOptions{
		URL: "https://github.com/brotherlogic/tasklister",
	})

	return err
}
