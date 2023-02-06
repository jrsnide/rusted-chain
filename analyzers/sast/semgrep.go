package sast

import (
	"log"
	"os/exec"
)

func Run(code_dir string) error {
	cmd := exec.Command("semgrep")
	cmd.Args = []string{
		"--config",
		"auto",
		code_dir,
	}

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	return nil
}
