package journalctl

import "github.com/carapace-sh/carapace"

// ActionTransports completes transports
//
//	audit (read from the kernel audit subsystem)
//	driver (internally generated messages)
func ActionTransports() carapace.Action {
	return carapace.ActionValuesDescribed(
		"audit", "read from the kernel audit subsystem",
		"driver", "internally generated messages",
		"syslog", "received via the local syslog socket with the syslog protocol",
		"journal", "received via the native journal protocol",
		"stdout", "read from a service's standard output or",
		"kernel", "read from the kernel",
	).Tag("transports")
}
