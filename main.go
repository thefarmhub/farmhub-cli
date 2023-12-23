package main

import (
	"os"

	"github.com/thefarmhub/farmhub-cli/cmd/core"
)

func main() {
	cmd := core.NewCommand()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
