package fastfetch

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPresets completes presets
func ActionPresets() carapace.Action {
	return carapace.ActionExecCommand("fastfetch", "--list-presets", "autocompletion")(func(output []byte) carapace.Action {
		presets := strings.Split(strings.TrimRight(string(output), "\n"), "\n")
		return carapace.ActionValues(presets...)
	})
}
