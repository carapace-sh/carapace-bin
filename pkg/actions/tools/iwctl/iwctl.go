package iwctl

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionAdapters completes IWD adapter names.
func ActionAdapters() carapace.Action {
	return actionTableNames("adapter", "list")
}

// ActionDevices completes IWD wireless device names.
func ActionDevices() carapace.Action {
	return actionTableNames("device", "list")
}

// ActionKnownNetworks completes saved network names.
func ActionKnownNetworks() carapace.Action {
	return actionTableNames("known-networks", "list")
}

// ActionNetworks completes networks visible to a station.
func ActionNetworks(device string) carapace.Action {
	if device == "" {
		return carapace.ActionValues()
	}
	return actionTableNames("station", device, "get-networks")
}

// ActionBSSes completes BSS addresses visible to a station.
func ActionBSSes(device string) carapace.Action {
	if device == "" {
		return carapace.ActionValues()
	}
	return actionTableNames("station", device, "get-bsses")
}

// ActionSecurityTypes completes network security type arguments.
func ActionSecurityTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"open", "open network",
		"psk", "pre-shared key network",
		"8021x", "802.1X network",
	)
}

// ActionDeviceProperties completes writable device properties.
func ActionDeviceProperties() carapace.Action {
	return carapace.ActionValuesDescribed(
		"Mode", "operation mode",
		"Powered", "powered state",
	)
}

// ActionAdapterProperties completes writable adapter properties.
func ActionAdapterProperties() carapace.Action {
	return carapace.ActionValuesDescribed(
		"Powered", "powered state",
	)
}

// ActionKnownNetworkProperties completes writable known-network properties.
func ActionKnownNetworkProperties() carapace.Action {
	return carapace.ActionValuesDescribed(
		"AutoConnect", "automatic connection state",
	)
}

// ActionPropertyValues completes values for writable properties.
func ActionPropertyValues(property string) carapace.Action {
	switch property {
	case "AutoConnect":
		return carapace.ActionValues("yes", "no")
	case "Mode":
		return carapace.ActionValues("ad-hoc", "ap", "station")
	case "Powered":
		return carapace.ActionValues("on", "off")
	default:
		return carapace.ActionValues()
	}
}

func actionTableNames(args ...string) carapace.Action {
	return carapace.ActionExecCommand("iwctl", args...)(func(output []byte) carapace.Action {
		return carapace.ActionValues(parseTableNames(output)...)
	})
}

func parseTableNames(output []byte) []string {
	lines := strings.Split(string(output), "\n")
	names := make([]string, 0)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.Contains(line, "---") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) == 0 || skipFirstField(fields[0]) {
			continue
		}

		names = append(names, fields[0])
	}

	return names
}

func skipFirstField(field string) bool {
	switch field {
	case "Name", "Address", "Adapter", "Adapters", "Available", "Devices", "Known", "Networks":
		return true
	default:
		return strings.HasPrefix(field, "[iwd]")
	}
}
