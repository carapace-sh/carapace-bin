package tsh

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionScopes completes scopes at which the user has assigned privileges
//
//	/example
//	/example/foo
func ActionScopes() carapace.Action {
	return carapace.ActionExecCommand("tsh", "scopes", "ls")(func(output []byte) carapace.Action {
		scopes := make([]string, 0)
		for line := range strings.SplitSeq(string(output), "\n") {
			if line = strings.TrimSpace(line); line != "" {
				scopes = append(scopes, line)
			}
		}
		return carapace.ActionValues(scopes...)
	})
}
