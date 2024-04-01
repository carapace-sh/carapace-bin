package xclip

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionTargets completes selection targets
func ActionTargets() carapace.Action {
	return carapace.ActionExecCommand("xclip", "-t", "TARGETS", "-o")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		values := make([]string, 0)
		for _, line := range lines {
			if line != "" {
				values = append(values, line)
			}
		}
		return carapace.ActionValues(values...)
	})
}
