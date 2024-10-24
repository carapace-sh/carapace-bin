package main

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd"
)

var commit, date string
var version = "develop"

func main() {
	if strings.HasSuffix(version, "-next") {
		version += fmt.Sprintf(" (%v) [%v]", date, commit)
	}
	cmd.Execute(version)
}
