package adb

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionDevices completes adb device serial numbers
//
// Device states from https://android.googlesource.com/platform/packages/modules/adb/+/refs/heads/main/adb.cpp
func ActionDevices() carapace.Action {
	return carapace.ActionExecCommand("adb", "devices", "-l")(func(output []byte) carapace.Action {
		vals := make([]string, 0)

		validStates := map[string]bool{
			"device":       true,
			"offline":      true,
			"bootloader":   true,
			"recovery":     true,
			"rescue":       true,
			"sideload":     true,
			"unauthorized": true,
			"authorizing":  true,
			"connecting":   true,
			"detached":     true,
		}

		for _, line := range strings.Split(string(output), "\n") {
			fields := strings.Fields(line)
			if len(fields) < 2 || !validStates[fields[1]] {
				continue
			}

			serial := fields[0]
			state := fields[1]
			props := make(map[string]string)
			for _, field := range fields[2:] {
				if key, value, ok := strings.Cut(field, ":"); ok {
					props[key] = value
				}
			}

			var desc string
			if model := props["model"]; model != "" {
				desc = strings.ReplaceAll(model, "_", " ")
				desc = strings.ReplaceAll(desc, "x86 64", "x86_64")
			}
			if state != "device" {
				if desc != "" {
					desc += " "
				}
				desc += "[" + state + "]"
			}
			if usb := props["usb"]; usb != "" {
				if desc != "" {
					desc += " "
				}
				desc += "(usb:" + usb + ")"
			}

			vals = append(vals, serial, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
