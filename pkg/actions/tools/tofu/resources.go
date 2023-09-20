package tofu

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionResources completes resources
func ActionResources(state string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"state", "list"}
		if state != "" {
			args = append(args, "--state", state)
		}
		return carapace.ActionExecCommand("tofu", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
