package main

import (
	"cacing/interface/cli"
	"log"
	"os"
)

func main() {
	err := cli.NewCliApp(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
