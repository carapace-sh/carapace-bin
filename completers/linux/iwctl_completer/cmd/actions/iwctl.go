package actions

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionStations() carapace.Action {
	return actionFromList("station", "list")
}

func ActionDevices() carapace.Action {
	return actionFromList("device", "list")
}

func ActionAdapters() carapace.Action {
	return actionFromList("adapter", "list")
}

func ActionKnownNetworks() carapace.Action {
	return actionFromList("known-networks", "list")
}

func ActionNetworks(station string) carapace.Action {
	if station == "" {
		return carapace.ActionValues()
	}
	return carapace.ActionExecCommand("iwctl", "station", station, "get-networks")(func(output []byte) carapace.Action {
		return valuesFromTable(output)
	})
}

func actionFromList(args ...string) carapace.Action {
	return carapace.ActionExecCommand("iwctl", args...)(func(output []byte) carapace.Action {
		return valuesFromTable(output)
	})
}

func valuesFromTable(output []byte) carapace.Action {
	lines := strings.Split(string(output), "\n")
	values := make([]string, 0, len(lines)*2)
	for _, line := range lines {
		line = strings.TrimSpace(strings.Trim(line, "-* "))
		fields := strings.Fields(line)
		if len(fields) == 0 || strings.HasPrefix(fields[0], "Name") || strings.HasPrefix(fields[0], "---") {
			continue
		}
		values = append(values, fields[0], strings.Join(fields[1:], " "))
	}
	return carapace.ActionValuesDescribed(values...)
}
