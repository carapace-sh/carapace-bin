package simctl

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
)

type simctlOutput struct {
	Devices map[string][]struct {
		State string `json:"state"`
		Name  string `json:"name"`
		UDID  string `json:"udid"`
	} `json:"devices"`
}

// ActionDevices completes iOS Simulator device UUIDs and names
//
//	00000000-0000-0000-0000-000000000000 (iPhone 14)
//	11111111-1111-1111-1111-111111111111 (iPad Pro)
func ActionDevices() carapace.Action {
	return carapace.ActionExecCommand("xcrun", "simctl", "list", "devices", "-j")(func(output []byte) carapace.Action {
		var result simctlOutput
		if err := json.NewDecoder(bytes.NewReader(output)).Decode(&result); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, devices := range result.Devices {
			for _, device := range devices {
				if device.UDID != "" {
					name := device.Name
					if name == "" {
						name = device.UDID
					}
					vals = append(vals, device.UDID, name)
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("simulator devices").Uid("simctl", "devices")
}

// ActionDevicesByState completes iOS Simulator device UUIDs filtered by state
//
//	00000000-0000-0000-0000-000000000000 (Booted)
//	11111111-1111-1111-1111-111111111111 (Shutdown)
func ActionDevicesByState(state string) carapace.Action {
	return carapace.ActionExecCommand("xcrun", "simctl", "list", "devices", "-j")(func(output []byte) carapace.Action {
		var result simctlOutput
		if err := json.NewDecoder(bytes.NewReader(output)).Decode(&result); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, devices := range result.Devices {
			for _, device := range devices {
				if device.UDID != "" && strings.EqualFold(device.State, state) {
					name := device.Name
					if name == "" {
						name = device.UDID
					}
					vals = append(vals, device.UDID, name)
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("simulator devices").Uid("simctl", "devices")
}
