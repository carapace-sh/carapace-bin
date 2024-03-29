package bluetoothctl

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionControllers completes controllers
func ActionControllers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("bluetoothctl", "list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines {
				if line == "" {
					break
				}

				fields := strings.Fields(line)
				vals = append(vals, fields[1], strings.Join(fields[2:], " "))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("bluetooth controllers")
}
