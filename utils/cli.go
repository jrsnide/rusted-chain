package utils

import (
	"flag"
)

func CliArgs() string {
	github_repo := flag.String("repo", "", "Github repository to analyze")

	flag.Parse()

	return *github_repo
}
