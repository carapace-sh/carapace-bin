package security

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionKeychains completes keychain names
//
//	Login.keychain
//	System.keychain
func ActionKeychains() carapace.Action {
	return carapace.ActionExecCommand("security", "list-keychains")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for line := range strings.SplitSeq(string(output), "\n") {
			path := strings.TrimSpace(strings.Trim(line, "\""))
			if path != "" {
				name := path
				if idx := strings.LastIndex(path, "/"); idx != -1 {
					name = path[idx+1:]
				}
				vals = append(vals, name, path)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("keychains").Uid("tools.security", "keychains")
}
