package bazel

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionTargets completes targets
//
//	//:windows_arm64
//	//docs:docs
func ActionTargets() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("bazel", "query", "//...")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines[:len(lines)-1] {
				vals = append(vals, strings.TrimPrefix(line, "//"))
			}
			return carapace.ActionValues(vals...)
		}).MultiParts("/", ":").Prefix("//")
	}).Tag("targets")
}
