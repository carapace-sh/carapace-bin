package herdr

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionIntegrations completes integrations
//
//	pi
//	omp
func ActionIntegrations() carapace.Action {
	return carapace.ActionExecCommand("herdr", "integration", "status")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for line := range strings.SplitSeq(string(output), "\n") {
			if line == "" {
				continue
			}
			if name := strings.SplitN(line, ":", 2); len(name) > 0 {
				vals = append(vals, name[0])
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("integrations")
}
