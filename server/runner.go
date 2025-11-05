package server

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

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

func (s *Server) Test(keyFromEnv string) error {
	key := strings.Replace(keyFromEnv, "\\n", "\n", -1)

	// Username must be "git" for SSH auth to work, not your real username.
	// See https://github.com/src-d/go-git/issues/637
	publicKey, err := ssh.NewPublicKeys("git", []byte(key), "")
	if err != nil {
		log.Fatalf("creating ssh auth method: %v -> %v", err, key)
	}

	// Get a tmp directory
	log.Printf("Checking out")
	tDir, err := os.MkdirTemp("", "gitcheckout")
	if err != nil {
		return err
	}

	t1 := time.Now()
	repo, err := git.PlainClone(tDir, false, &git.CloneOptions{
		Auth: publicKey,
		URL:  "git@github.com:brotherlogic/tasklister.git",
	})
	log.Printf("Took %v to clone", time.Since(t1))
	if err != nil {
		return fmt.Errorf("error in clone: %w", err)
	}

	// Update a file
	err = writeString(fmt.Sprintf("%v/%v", tDir, "test.txt"), "Hello\n")
	if err != nil {
		return fmt.Errorf("error writing string: %w", err)
	}

	// Commit the change
	w, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("error in worktree: %w", err)
	}
	_, err = w.Add(fmt.Sprintf("test.txt"))
	if err != nil {
		return fmt.Errorf("error in add: %w", err)
	}

	_, err = w.Commit("Updating tasklist", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "tasklister",
			Email: "tasklister@brotherlogic.com",
		}})
	if err != nil {
		return fmt.Errorf("error in commit: %w", err)
	}

	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       publicKey,
	})
	if err != nil {
		return fmt.Errorf("Error in push: %w", err)
	}
	return nil
}
