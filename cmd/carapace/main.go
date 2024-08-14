package main

import (
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd"
)

var version = "develop"

func main() {
	cmd.Execute(version)
}
