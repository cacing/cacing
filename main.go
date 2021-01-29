package main

import (
	"fmt"
	"os"

	"github.com/cacing/cacing/interface/cli"
)

func main() {
	err := cli.NewCliApp(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
