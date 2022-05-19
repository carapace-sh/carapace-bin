package tmux

import "github.com/rsteube/carapace"

// ActionClientFlags completes client flags
//   read-only (the client is read-only)
//   wait-exit (wait for an empty line input before exiting in control mode)
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
