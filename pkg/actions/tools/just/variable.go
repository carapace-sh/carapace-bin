package just

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionVariables completes variables
// variablea
// variableb
func ActionVariables(justfile string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"--variables"}
		if justfile != "" {
			args = append(args, "--justfile", justfile)
		}

		return carapace.ActionExecCommand("just", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(strings.Fields(lines[0])...)
		})
	})
}
