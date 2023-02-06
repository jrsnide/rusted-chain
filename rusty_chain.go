package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/rusty-chain/utils"
	"github.com/urfave/cli"
)

var tool_analysis_dir string

func init() {
	home_dir, dir_err := os.UserHomeDir()

	if dir_err == nil {
		tool_analysis_dir = home_dir + "/.rusty_chain"
	} else {
		fmt.Println("There was an issue getting the user's home directory")
		log.Fatal(dir_err)
	}
}
func main() {
	log.SetPrefix("rusty-chain: ")

	app, err := utils.CliArgs()

	if err != nil {
		log.Fatal("error")
	}
	app.Action = func(c *cli.Context) error {

		//check if analysis dir exists
		_, exists_err := os.Stat(tool_analysis_dir)

		if exists_err == nil {
			create_dir_err := os.Mkdir(tool_analysis_dir, fs.ModeDir)

			if create_dir_err != nil {
				fmt.Printf("There was an error creating %v\n", tool_analysis_dir)
				log.Fatal(create_dir_err)
			}
		}

		repo_url := c.String("repo")
		fmt.Printf("Cloning repo %v\n", repo_url)
		clone_err := utils.Clone(repo_url, tool_analysis_dir)

		if clone_err != nil {
			log.Fatal(clone_err)
		}

		fmt.Printf("Running SAST analysis against %v\n", c.String("repo"))

		return nil
	}

	run_err := app.Run(os.Args)
	if run_err != nil {
		fmt.Println(run_err)
		empty_args := make([]string, 1)
		app.Run(empty_args)
		cli.OsExiter(1)
	}

}
