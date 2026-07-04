package os

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

func groupStyle(id string) string {
	if id == "0" {
		return style.Red
	}
	if _id, err := strconv.Atoi(id); err == nil && _id >= 1000 {
		return style.Blue
	}
	return style.Default
}

// ActionGroups completes system group names
//
//	root (0)
//	ssh (101)
func ActionGroups() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		groups := []string{}
		seen := make(map[string]bool)

		if content, err := os.ReadFile("/etc/group"); err == nil {
			for entry := range strings.SplitSeq(string(content), "\n") {
				splitted := strings.Split(entry, ":")
				if len(splitted) > 2 {
					group := splitted[0]
					id := splitted[2]

					if len(strings.TrimSpace(group)) > 0 {
						groups = append(groups, group, id, groupStyle(id))
						seen[group] = true
					}
				}
			}
		}

		if runtime.GOOS == "darwin" {
			if output, err := c.Command("dscl", ".", "-list", "/Groups", "PrimaryGroupID").Output(); err == nil {
				for line := range strings.SplitSeq(string(output), "\n") {
					fields := strings.Fields(line)
					if len(fields) == 2 {
						group, id := fields[0], fields[1]
						if !seen[group] && len(strings.TrimSpace(group)) > 0 {
							groups = append(groups, group, id, groupStyle(id))
						}
					}
				}
			}
		}

		return carapace.ActionStyledValuesDescribed(groups...)
	}).Tag("groups").Uid("os", "group")
}

// ActionGroupMembers completes system group members
//
//	root
//	daemon
func ActionGroupMembers(group string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := []string{}
		if content, err := os.ReadFile("/etc/group"); err == nil {
			for entry := range strings.SplitSeq(string(content), "\n") {
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

		if runtime.GOOS == "darwin" {
			if output, err := c.Command("dscl", ".", "-read", "/Groups/"+group, "GroupMembership").Output(); err == nil {
				line := strings.TrimSpace(string(output))
				if prefix, members, ok := strings.Cut(line, ":"); ok && strings.TrimSpace(prefix) == "GroupMembership" {
					for _, member := range strings.Fields(strings.TrimSpace(members)) {
						vals = append(vals, member)
					}
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}
