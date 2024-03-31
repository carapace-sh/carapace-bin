package xclip

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionTarget completes selection targets
func ActionTarget() carapace.Action {
	return carapace.ActionExecCommand("xclip", "-t", "TARGETS", "-o")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		values := []string{"TARGETS"}
		for _, line := range lines {
			if line != "" {
				values = append(values, line)
			}
		}
		return carapace.ActionValues(values...)
	})
}
