package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionDrivers completes molecule drivers
func ActionDrivers() carapace.Action {
	return carapace.ActionExecCommand("molecule", "drivers", "-f", "plain")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		drivers := make([]string, 0)
		for _, driver := range lines {
			if driver != "" {
				drivers = append(drivers, driver)
			}
		}
		return carapace.ActionValues(drivers...)
	})
}
