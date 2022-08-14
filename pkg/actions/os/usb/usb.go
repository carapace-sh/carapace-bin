// package usb contains usb related actions
package usb

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionDeviceNumbers completes usb devices
//
//	001:006 (Apple, Inc. Bluetooth USB Host Controller)
//	002:001 (Linux Foundation 3.0 root hub)
func ActionDeviceNumbers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("lsusb")(func(output []byte) carapace.Action {
			r := regexp.MustCompile(`^Bus (?P<bus>\d+) Device (?P<device>\d+): ID (?P<vendor>[^ ]+):(?P<product>[^ ]+) (?P<name>.*)$`)

			vals := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				if r.MatchString(line) {
					matches := r.FindStringSubmatch(line)
					vals = append(vals, fmt.Sprintf("%v:%v", matches[1], matches[2]), matches[5])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

// ActionProductNumbers completes usb products
//
//	0a5c:4500 (Broadcom Corp. BCM2046B1 USB 2.0 Hub (part of BCM2046 Bluetooth))
//	1d6b:0002 (Linux Foundation 2.0 root hub)
func ActionProductNumbers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("lsusb")(func(output []byte) carapace.Action {
			r := regexp.MustCompile(`^Bus (?P<bus>\d+) Device (?P<device>\d+): ID (?P<vendor>[^ ]+):(?P<product>[^ ]+) (?P<name>.*)$`)

			vals := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				if r.MatchString(line) {
					matches := r.FindStringSubmatch(line)
					vals = append(vals, fmt.Sprintf("%v:%v", matches[3], matches[4]), matches[5])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
