package adb

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionDevices completes adb device serial numbers
func ActionDevices() carapace.Action {
	return carapace.ActionExecCommand("adb", "devices", "-l")(func(output []byte) carapace.Action {
		vals := make([]string, 0)

		for _, line := range strings.Split(string(output), "\n") {
			fields := strings.Fields(line)
			if len(fields) < 2 || fields[1] != "device" {
				continue
			}

			serial := fields[0]
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
				if usb := props["usb"]; usb != "" {
					desc += " (usb:" + usb + ")"
				}
			}

			vals = append(vals, serial, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
