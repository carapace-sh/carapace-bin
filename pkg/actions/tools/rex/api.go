package rex

import (
	"strings"
	"time"

	"github.com/carapace-sh/carapace"
)

// ActionApiMethods completes server API method names
//
//	session.list
//	block.close
func ActionApiMethods() carapace.Action {
	return carapace.ActionExecCommand("rex", "api", "list", "-q")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, line := range strings.Split(string(output), "\n") {
			if fields := strings.Fields(line); len(fields) > 0 {
				vals = append(vals, fields[0])
			}
		}
		return carapace.ActionValues(vals...)
	}).Cache(24 * time.Hour).Tag("methods")
}
