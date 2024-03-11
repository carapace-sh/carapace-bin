package golang

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPackages completes packages
//
//	encoding/json
//	github.com/carapace-sh/carapace-bin/completers/go_completer
func ActionPackages() carapace.Action {
	return carapace.ActionExecCommand("go", "list", "-f", "{{.ImportPath}}\n{{.Doc}} ", "all")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
	})
}
