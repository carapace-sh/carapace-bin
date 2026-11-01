package action

import "github.com/carapace-sh/carapace"

func ActionFacilities() carapace.Action {
	return carapace.ActionValuesDescribed(
		"auth", "security/authorization messages",
		"daemon", "system daemons",
		"kern", "kernel messages",
		"lpr", "line printer subsystem",
		"mail", "mail system",
		"news", "network news subsystem",
		"syslog", "messages generated internally by syslogd",
		"user", "random user-level messages",
	)
}
