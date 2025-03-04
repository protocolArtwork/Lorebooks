package newdb

import (
	"flag"
	"fmt"
	"log"
)

func SubCommand() {
	name := flag.String("name", "World", "the name to use in the greeting")
	flag.Parse()

	fmt.Printf("Hello, %s!\n", *name)

	if flag.NArg() > 0 {
		log.Println("Additional arguments:", flag.Args())
	}
}
