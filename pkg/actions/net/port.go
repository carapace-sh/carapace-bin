package net

import (
	"github.com/rsteube/carapace"
)

// ActionPorts completes commonly used ports (opiniated)
//
//	80 (http)
//	443 (https)
func ActionPorts() carapace.Action {
	return carapace.ActionValuesDescribed(
		"21", "ftp",
		"22", "ssh",
		"25", "smtp",
		"80", "http",
		"110", "pop3",
		"143", "imap",
		"443", "https",
		"587", "smtp ssl",
		"993", "imap ssl",
		"995", "pop3 ssl",
		"2077", "webdav",
		"2078", "webdav ssl",
		"2082", "cpanel",
		"2083", "cpanel ssl",
		"2086", "webhost manager",
		"2087", "webhost manager ssl",
		"2095", "webmail",
		"2096", "webmail ssl",
		"3306", "mysql",
		"8080", "http alternative",
		"8443", "https alternative",
	)
}
