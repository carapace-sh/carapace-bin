package bazel

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionInfoKeys completes info keys
//
//	announce_rc (Announce rc file options)
//	build_path (Build path)
func ActionInfoKeys() carapace.Action {
	return carapace.ActionExecCommand("bazel", "help", "info-keys")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if line == "" {
				continue
			}
			vals = append(vals, strings.Fields(line)[0])
		}
		return carapace.ActionValues(vals...)
	}).Tag("info keys")
}