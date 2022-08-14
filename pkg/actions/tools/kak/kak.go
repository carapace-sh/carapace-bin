// package kak contains kakoune related actions
package kak

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionSessions completes kak sessions
//
//	12345
//	some_name
func ActionSessions() carapace.Action {
	return carapace.ActionExecCommand("kak", "-l")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
