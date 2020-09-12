package action

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionYesNo() carapace.Action {
	return carapace.ActionValues("yes", "no")
}

func ActionConnections() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("nmcli", "--terse", "connection", "show").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, (len(lines)-1)*2)
			for index, line := range lines[:len(lines)-1] {
				parts := strings.Split(line, ":")
				vals[index*2] = parts[0]
				vals[index*2+1] = parts[1]
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}

// ActionDevices completes network devices (optionally filtered by deviceType)
//   ActionDevices("") // all
//   ActionDevices("wifi") // only wifi devices
func ActionDevices(deviceType string) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("nmcli", "--terse", "--fields", "device,type", "device", "status").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				parts := strings.Split(line, ":")
				if deviceType == "" || deviceType == parts[1] {
					vals = append(vals, parts[0], parts[1])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}

func ActionBssids() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("nmcli", "--terse", "--fields", "bssid,ssid", "device", "wifi", "list").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, (len(lines)-1)*2)
			for index, line := range lines[:len(lines)-1] {
				vals[index*2] = line[:22]
				vals[index*2+1] = line[23:]
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}

func ActionSsids() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("nmcli", "--terse", "--fields", "bssid,ssid", "device", "wifi", "list").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				if ssid := line[23:]; ssid != "" {
					vals = append(vals, line[23:], line[:22])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}
