package tmux

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionClientFlags completes client flags
//
//	read-only (the client is read-only)
//	wait-exit (wait for an empty line input before exiting in control mode)
func ActionClientFlags() carapace.Action {
	return carapace.ActionValuesDescribed(
		"active-pane", "the client has an independent active pane",
		"ignore-size", "the client does not affect the size of other clients",
		"no-output", "the client does not receive pane output in control mode",
		"pause-after=", "output is paused once the pane is seconds behind in control mode",
		"read-only", "the client is read-only",
		"wait-exit", "wait for an empty line input before exiting in control mode",
	)
}

// ActionClients completes clients
//
//	/dev/pts/0: 0 (alacritty)
//	/dev/pts/1: 1 (kitty)
func ActionClients() carapace.Action {
	return carapace.ActionExecCommand("tmux", "list-clients")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if line == "" {
				continue
			}
			if splitted := strings.SplitN(line, ": ", 2); len(splitted) == 2 {
				vals = append(vals, splitted...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
