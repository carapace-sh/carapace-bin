package saw

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionStreams completes log streams.
//
//	streamOne
//	streamTwo
func ActionStreams(group string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("saw", "streams", "--prefix", c.Value, group)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines...)
		})
	}).Tag("streams")
}
