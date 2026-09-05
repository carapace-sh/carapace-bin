package bazel

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionInfoKeys completes info keys
//
//	bazel-bin (Configuration dependent directory for binaries)
//	client-env (The specifications that need to be added to the project-specific rc file to freeze the current client environment)
func ActionInfoKeys() carapace.Action {
	return carapace.ActionExecCommand("bazel", "help", "info-keys")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if line == "" {
				continue
			}
			fields := strings.Fields(line)
			vals = append(vals, fields[0], strings.Join(fields[1:], " "))
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("info keys")
}
