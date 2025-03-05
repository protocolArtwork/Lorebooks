package main

import (
	"flag"

	"./lbcli"
	"./newdb"
)

func main() {
	flag.Parse()
	var args []string = flag.Args()

	if len(args) < 1 {
		lbcli.Fatal("No database name provided")
		return
	} else if len(args) < 2 {
		lbcli.Fatal("No subcommand provided")
		return
	}

	// subcommand name
	var subcommand string
	{
		subcommand = args[0]
	}

	switch subcommand {
	case "newdb":
		dbrelpath := args[1]
		newdb.NewDB(dbrelpath)
		break
	default:
		lbcli.Fatal("Unknown subcommand: " + args[1])
	}
}
