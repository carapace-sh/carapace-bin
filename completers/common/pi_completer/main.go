package main

import (
	"fmt"
	"os"

	"github.com/carapace-sh/carapace-bin/completers/common/pi_completer/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}