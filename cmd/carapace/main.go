package main

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd"
)

var commit, date string
var version = "develop"

//go:generate sh -c "rm -rf ../../completers_release && go run ../carapace-generate copy ../../completers"
//go:generate sh -c "go run ../carapace-generate completers --code --tag !release ../../completers darwin          > cmd/completers/completers_darwin.go"
//go:generate sh -c "go run ../carapace-generate completers --code --tag !release ../../completers linux           > cmd/completers/completers_linux.go"
//go:generate sh -c "go run ../carapace-generate completers --code --tag !release ../../completers windows         > cmd/completers/completers_windows.go"
//go:generate sh -c "go run ../carapace-generate completers --code --tag  release ../../completers_release darwin  > cmd/completers/completers_release_darwin.go"
//go:generate sh -c "go run ../carapace-generate completers --code --tag  release ../../completers_release linux   > cmd/completers/completers_release_linux.go"
//go:generate sh -c "go run ../carapace-generate completers --code --tag  release ../../completers_release windows > cmd/completers/completers_release_windows.go"
//go:generate sh -c "go run ../carapace-generate macros --code github.com/carapace-sh/carapace-bin/pkg/actions github.com/carapace-sh/carapace-bridge/pkg/actions --output ../../pkg/actions/actions_generated.go"
//go:generate sh -c "go run ../carapace-generate conditions ../../pkg/conditions --output ../../pkg/conditions/conditions_generated.go"
func main() {
	if strings.Contains(version, "SNAPSHOT") {
		version += fmt.Sprintf(" (%v) [%v]", date, commit)
	}
	cmd.Execute(version)
}
