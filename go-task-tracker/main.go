package main

import (
	"log"
	"os"
	"task-cli/cli"
)

const timeFormat = "2006/01/02 15:04"

// TODO: implement func getLastID in json_repository
// TODO: implement func AddTask in service.go

func main() {
	args := os.Args[1:]

	if err := cli.Run(args); err != nil {
		log.Fatal(err)
	}
}
