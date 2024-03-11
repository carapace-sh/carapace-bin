package os

import (
	"os"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionGroups completes system group names
//
//	root (0)
//	ssh (101)
func ActionGroups() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		groups := []string{}
		if content, err := os.ReadFile("/etc/group"); err == nil {
			for _, entry := range strings.Split(string(content), "\n") {
				splitted := strings.Split(entry, ":")
				if len(splitted) > 2 {
					group := splitted[0]
					id := splitted[2]

					_style := style.Default
					if id == "0" {
						_style = style.Red
					} else if _id, err := strconv.Atoi(id); err == nil && _id >= 1000 {
						_style = style.Blue
					}

					if len(strings.TrimSpace(group)) > 0 {
						groups = append(groups, group, id, _style)
					}
				}
			}
		}
		return carapace.ActionStyledValuesDescribed(groups...)
	}).Tag("groups")
}

// ActionGroupMembers completes system group members
//
//	root
//	daemon
func ActionGroupMembers(group string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := []string{}
		if content, err := os.ReadFile("/etc/group"); err == nil {
			for _, entry := range strings.Split(string(content), "\n") {
				splitted := strings.Split(entry, ":")
				if len(splitted) > 3 &&
					splitted[0] == group {
					if len(strings.TrimSpace(group)) > 0 &&
						len(strings.TrimSpace(splitted[3])) > 0 {
						vals = append(vals, strings.Split(splitted[3], ",")...)
					}
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}
