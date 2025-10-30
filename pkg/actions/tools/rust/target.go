package rust

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionTargets completes targets
//
//	aarch64-apple-darwin
//	aarch64-apple-ios
func ActionTargets() carapace.Action {
	return carapace.ActionExecCommand("rustc", "--print", "target-list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines...)
	}).Tag("targets").Uid("rust", "target")
}
