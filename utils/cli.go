package utils

import (
	"errors"
	"log"
	"regexp"

	"github.com/urfave/cli"
)

var repo_regex *regexp.Regexp
var repo_regex_err error

func init() {
	repo_regex, repo_regex_err = regexp.Compile(`^https\:\/\/github\.com\/[a-zA-Z0-9_\-\/]*$`)

	if repo_regex_err != nil {
		log.Fatal(repo_regex_err)
	}
}

func validate_flag(regex *regexp.Regexp, flag string, error_msg string) error {
	flag_val_correct := regex.Match([]byte(flag))

	if flag_val_correct {
		return nil
	} else {
		return errors.New(error_msg)
	}
}

func CliArgs() (cli.App, error) {
	app := cli.NewApp()
	app.Name = "Rusty Chain"
	app.Usage = ""
	app.Description = "Supply chain security assessment tool for OSS packages"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "repo", Value: "", Usage: "Github URL to assess (Required)", Required: true},
		cli.StringFlag{Name: "version", Value: "", Usage: "OSS Version to assess (If blank, latest version will be used)"},
	}
	app.Before = func(ctx *cli.Context) error {
		repo_correct := validate_flag(repo_regex, ctx.String("repo"), "error validating provided Github repository in --repo")

		return repo_correct
	}

	app.UsageText = "rusty-chain --repo=https://github.com/lodash/lodash"
	app.Author = "Jesse Snider"
	app.Email = "theappsecguy@gmail.com"

	return *app, nil
}
