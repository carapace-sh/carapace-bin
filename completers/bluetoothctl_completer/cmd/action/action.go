package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionAgents() carapace.Action {
	return carapace.ActionValues(
		"off",
		"on",
		"auto",
		"NoInputNoOutput",
		"KeyboardOnly",
		"KeyboardDisplay",
		"DisplayYesNo",
		"DisplayOnly",
	)
}

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
	})
}

func ActionDevices(filter string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"devices"}
		if filter != "" {
			args = append(args, filter)
		}
		return carapace.ActionExecCommand("bluetoothctl", args...)(func(output []byte) carapace.Action {
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
	})
}
