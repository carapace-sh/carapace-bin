package net

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionPorts completes known and used ports.
func ActionPorts() carapace.Action {
	return carapace.Batch(
		ActionKnownPorts(),
		ActionUsedPorts(),
	).ToA()
}

// ActionKnownPorts completes commonly used ports (opiniated)
//
//	80 (http)
//	443 (https)
func ActionKnownPorts() carapace.Action {
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
	).Tag("known ports")
}

// ActionUsedPorts completes ports currently in use
//
//	3000 (used by: mdbook)
//	8000 (used by: python)
func ActionUsedPorts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO support other tools
		if _, err := exec.LookPath("lsof"); err == nil {
			return carapace.ActionExecCommand("lsof", "-i", "-P", "-n")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")

				vals := make([]string, 0)
				for _, line := range lines {
					if fields := strings.Fields(line); len(fields) > 9 {
						if fields[len(fields)-1] != "(LISTEN)" {
							continue
						}

						if splitted := strings.Split(fields[len(fields)-2], ":"); len(splitted) > 1 {
							vals = append(vals, splitted[len(splitted)-1], "used by: "+fields[0])
						}
					}
				}
				return carapace.ActionValuesDescribed(vals...).Style(style.Carapace.KeywordNegative)
			})
		}
		return carapace.ActionValues()
	}).Tag("used ports")
}
