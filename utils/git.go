package utils

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func Clone(repo_url string, clone_dir string) error {

	_, err := git.PlainClone(clone_dir, false, &git.CloneOptions{
		URL:      repo_url,
		Progress: os.Stdout,
	})

	return err
}
