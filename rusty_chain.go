package main

import (
	"fmt"

	"github.com/rusty-chain/utils"
)

func main() {
	repo := utils.CliArgs()
	fmt.Println(repo)
}
