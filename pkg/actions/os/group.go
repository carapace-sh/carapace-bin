package os

import (
	"io/ioutil"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionGroups completes system group names
//    root (0)
//    ssh (101)
func ActionGroups() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		groups := []string{}
		if content, err := ioutil.ReadFile("/etc/group"); err == nil {
			for _, entry := range strings.Split(string(content), "\n") {
				splitted := strings.Split(entry, ":")
				if len(splitted) > 2 {
					group := splitted[0]
					id := splitted[2]
					if len(strings.TrimSpace(group)) > 0 {
						groups = append(groups, group, id)
					}
				}
			}
		}
		return carapace.ActionValuesDescribed(groups...)
	})
}

// ActionGroupMembers completes system group members
//   root
//   daemon
func ActionGroupMembers(group string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := []string{}
		if content, err := ioutil.ReadFile("/etc/group"); err == nil {
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
