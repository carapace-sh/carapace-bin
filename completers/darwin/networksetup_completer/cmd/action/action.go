package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionNetworkServices() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("networksetup", "-listallnetworkservices")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			services := make([]string, 0)
			for _, line := range lines {
				if trimmed := strings.TrimSpace(line); trimmed != "" && !strings.HasPrefix(trimmed, "An asterisk") {
					services = append(services, trimmed)
				}
			}
			return carapace.ActionValues(services...)
		})
	}).Tag("network services")
}