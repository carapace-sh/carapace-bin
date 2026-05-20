package action

import (
	"os"
	"os/exec"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionAdapters completes wireless adapter names.
func ActionAdapters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if vals := iwctlColumn("adapter", "list"); len(vals) > 0 {
			return carapace.ActionValues(vals...)
		}
		return actionSysEntries("/sys/class/ieee80211")
	}).Tag("wireless adapters")
}

// ActionDevices completes wireless device names.
func ActionDevices() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if vals := iwctlColumn("device", "list"); len(vals) > 0 {
			return carapace.ActionValues(vals...)
		}
		return actionSysEntries("/sys/class/net")
	}).Tag("wireless devices")
}

// ActionKnownNetworks completes saved network names.
func ActionKnownNetworks() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionValues(iwctlColumn("known-networks", "list")...)
	}).Tag("known networks")
}

// ActionNetworks completes scanned networks for a station device.
func ActionNetworks(device string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionValues(iwctlColumn("station", device, "get-networks")...)
	}).Tag("wireless networks")
}

func actionSysEntries(path string) carapace.Action {
	entries, err := os.ReadDir(path)
	if err != nil {
		return carapace.ActionValues()
	}

	vals := make([]string, 0)
	for _, entry := range entries {
		vals = append(vals, entry.Name())
	}
	return carapace.ActionValues(vals...)
}

func iwctlColumn(args ...string) []string {
	output, err := exec.Command("iwctl", args...).Output()
	if err != nil {
		return nil
	}

	vals := make([]string, 0)
	for _, line := range strings.Split(string(output), "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" ||
			strings.HasPrefix(trimmed, "Name ") ||
			strings.HasPrefix(trimmed, "Devices") ||
			strings.HasPrefix(trimmed, "Adapters") ||
			strings.HasPrefix(trimmed, "Known Networks") ||
			strings.HasPrefix(trimmed, "Available networks") ||
			strings.HasPrefix(trimmed, "----") {
			continue
		}

		fields := strings.Fields(trimmed)
		if len(fields) > 0 {
			vals = append(vals, fields[0])
		}
	}
	return vals
}
