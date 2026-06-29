package main

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd"
)

var commit, date string
var version = "develop"

//go:generate go run ../carapace-generate generate-all ../../completers
//go:generate go run ../carapace-generate macros --code github.com/carapace-sh/carapace-bin/pkg/actions github.com/carapace-sh/carapace-bridge/pkg/actions github.com/carapace-sh/carapace-jjlex/pkg/actions --output ../../pkg/actions/actions_generated.go
//go:generate go run ../carapace-generate conditions ../../pkg/conditions --output ../../pkg/conditions/conditions_generated.go
func main() {
	if strings.Contains(version, "SNAPSHOT") {
		version += fmt.Sprintf(" (%v) [%v]", date, commit)
	}
	cmd.Execute(version)
}
