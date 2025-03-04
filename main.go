package main

import (
	"log"
	"newdb"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No subcommand provided")
	}

	switch os.Args[1] {
	case "newdb":
		newdb.SubCommand()
		break
	default:
		log.Fatal("Unknown subcommand: ", os.Args[1])
	}
}
