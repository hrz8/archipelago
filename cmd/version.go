package cmd

import (
	"flag"
	"fmt"

	"archipelago.dev/chat/config"
)

func VersionCmd() *flag.FlagSet {
	c := flag.NewFlagSet("serve", flag.ExitOnError)
	return c
}

func VersionHandler(c *flag.FlagSet) error {
	fmt.Println("Archipelago Chat version:", config.Version)
	return nil
}
