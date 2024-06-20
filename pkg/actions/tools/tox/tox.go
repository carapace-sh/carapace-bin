// package tox contains tox related actions
package tox

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionEnvironments completes tox environments
func ActionEnvironments() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		cmd := c.Command("tox", "--listenvs-all")
		if output, err := cmd.Output(); err != nil || cmd.ProcessState.ExitCode() != 0 {
			return carapace.ActionMessage(err.Error())
		} else if strings.Contains(string(output), "ROOT: No tox.ini") {
				return carapace.ActionValues()
		} else {
			lines := strings.Split(string(output), "\n")
			lines = append(lines, "ALL")
			return carapace.ActionValues(lines...)
		}
	})
}
