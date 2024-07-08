// package tox contains tox related actions
package tox

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionEnvironments completes tox environments
func ActionEnvironments() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("tox", "--listenvs-all")(func(output []byte) carapace.Action {
			if strings.Contains(string(output), "ROOT: No tox.ini") {
				return carapace.ActionValues()
			} else {
				lines := strings.Split(string(output), "\n")
				lines = append(lines, "ALL")
				return carapace.ActionValues(lines...)
			}
		})
	})
}
