package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mfranczy/compat/cmd/compat/validate"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		// TODO: add usage doc
		fmt.Println("Usage...")
		os.Exit(1)
	}

	cmd := args[0]
	cmdArgs := args[1:]

	switch cmd {
	case validate.CmdName:
		if err := validate.Run(cmdArgs); err != nil {
			log.Fatal(validate.CmdName, err)
		}
	default:
		log.Fatalf("Unknown command '%s'", cmd)
	}
}
