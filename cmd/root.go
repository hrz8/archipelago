package cmd

import (
	"errors"
	"os"
)

func Execute() error {
	if len(os.Args) < 2 {
		return errors.New("subcommand is required")
	}

	subCmd := os.Args[1]

	if subCmd == "serve" {
		return ServeHandler(ServeCmd())
	} else if subCmd == "version" {
		return VersionHandler(VersionCmd())
	}

	return errors.New("unknown subcommand")
}
