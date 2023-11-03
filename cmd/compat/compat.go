package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"compat/cmd/compat/validate"
	"compat/cmd/compat/version"
)

func main() {
	flag.Parse()
	args := flag.Args()

	// TODO: recognize if tool has been executed with root permissions
	if len(args) < 1 {
		// TODO: add usage doc
		fmt.Println("Usage...")
		os.Exit(1)
	}

	cmd := args[0]
	cmdArgs := args[1:]

	switch cmd {
	case validate.HostCmdName:
		if err := validate.RunHostCmd(cmdArgs); err != nil {
			log.Fatalf("%s command error: %s", validate.HostCmdName, err)
		}
		break
	case validate.SchemaCmdName:
		if err := validate.RunSchemaCmd(cmdArgs); err != nil {
			log.Fatalf("%s command error: %s", validate.SchemaCmdName, err)
		}
		break
	case version.CmdName:
		version.Run()
		break
	default:
		log.Fatalf("Unknown command '%s'", cmd)
	}
}
