package brew

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionInstalledTaps completes installed taps
//
//	homebrew/cask
//	homebrew/core
func ActionInstalledTaps() carapace.Action {
	return carapace.ActionExecCommand("brew", "tap")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	}).Tag("taps")
}
