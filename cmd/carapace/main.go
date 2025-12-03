package main

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd"
)

var commit, date string
var version = "develop"

//go:generate sh -c "rm -rf ../../completers_release && go run ../carapace-generate copy ../../completers"
//go:generate sh -c "go run ../carapace-generate completers --code --tag '!release && !force_all' ../../completers darwin            --output cmd/completers/completers_darwin.go"
//go:generate sh -c "go run ../carapace-generate completers --code --tag '!release && !force_all' ../../completers linux             --output cmd/completers/completers_linux.go"
//go:generate sh -c "go run ../carapace-generate completers --code --tag '!release && !force_all' ../../completers windows           --output cmd/completers/completers_windows.go"
//go:generate sh -c "go run ../carapace-generate completers --code --tag '!release &&  force_all' ../../completers force_all         --output cmd/completers/completers_all.go"
//go:generate sh -c "go run ../carapace-generate completers --code --tag  'release && !force_all' ../../completers_release darwin    --output cmd/completers/completers_release_darwin.go"
//go:generate sh -c "go run ../carapace-generate completers --code --tag  'release && !force_all' ../../completers_release linux     --output cmd/completers/completers_release_linux.go"
//go:generate sh -c "go run ../carapace-generate completers --code --tag  'release && !force_all' ../../completers_release windows   --output cmd/completers/completers_release_windows.go"
//go:generate sh -c "go run ../carapace-generate completers --code --tag  'release &&  force_all' ../../completers_release force_all --output cmd/completers/completers_release_all.go"
//go:generate sh -c "go run ../carapace-generate macros --code github.com/carapace-sh/carapace-bin/pkg/actions github.com/carapace-sh/carapace-bridge/pkg/actions --output ../../pkg/actions/actions_generated.go"
//go:generate sh -c "go run ../carapace-generate conditions ../../pkg/conditions --output ../../pkg/conditions/conditions_generated.go"
func main() {
	if strings.Contains(version, "SNAPSHOT") {
		version += fmt.Sprintf(" (%v) [%v]", date, commit)
	}
	cmd.Execute(version)
}
