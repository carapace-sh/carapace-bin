package bluetoothctl

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type DeviceOpts struct {
	Unknown   bool
	Paired    bool
	Bonded    bool
	Trusted   bool
	Connected bool
}

func (o DeviceOpts) Default() DeviceOpts {
	o.Unknown = true
	o.Paired = true
	o.Bonded = true
	o.Trusted = true
	o.Connected = true
	return o
}

// ActionDevices completes devices
//
//	00:11:22:33:44:AA (DeviceA)
//	00:11:22:33:44:BB (DeviceB)
func ActionDevices(opts DeviceOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch()
		if opts.Unknown {
			batch = append(batch, actionDevices(""))
		}
		if opts.Paired {
			batch = append(batch, actionDevices("Paired").Style(style.Yellow))
		}
		if opts.Bonded {
			batch = append(batch, actionDevices("Bonded").Style(style.Magenta))
		}
		if opts.Trusted {
			batch = append(batch, actionDevices("Trusted").Style(style.Green))
		}
		if opts.Connected {
			batch = append(batch, actionDevices("Connected").Style(style.Blue))
		}
		return batch.ToA()
	}).Tag("bluetooth devices")
}

func actionDevices(filter string) carapace.Action {
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
