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
		switch runtime.GOOS {
		case "windows":
			return actionGroupsWindows()
		default:
			return actionGroupsUnix()
		}
	}).Tag("groups").Uid("os", "group")
}

func actionGroupsUnix() carapace.Action {
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
	})
}

func actionGroupsWindows() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		groups := make([]string, 0)
		for _, cmd := range [][]string{
			{"net", "localgroup"},
			{"net", "group"},
		} {
			if output, err := c.Command(cmd[0], cmd[1:]...).Output(); err == nil {
				lines := strings.Split(strings.ReplaceAll(string(output), "\r", ""), "\n")
				skipHeader := true
				for _, line := range lines {
					trimmed := strings.TrimSpace(line)
					if trimmed == "" {
						skipHeader = false
						continue
					}
					if skipHeader {
						continue
					}
					if strings.HasPrefix(trimmed, "The command") || strings.HasPrefix(trimmed, "*") {
						continue
					}
					for _, group := range strings.Fields(line) {
						if len(group) > 0 {
							groups = append(groups, group, group, style.Blue)
						}
					}
				}
			}
		}
		return carapace.ActionStyledValuesDescribed(groups...)
	})
}

// ActionGroupMembers completes system group members
//
//	root
//	daemon
func ActionGroupMembers(group string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch runtime.GOOS {
		case "windows":
			return actionGroupMembersWindows(c, group)
		default:
			return actionGroupMembersUnix(group)
		}
	})
}

func actionGroupMembersUnix(group string) carapace.Action {
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
					vals = append(vals, strings.Fields(strings.TrimSpace(members))...)
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}

func actionGroupMembersWindows(c carapace.Context, group string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if output, err := c.Command("net", "localgroup", group).Output(); err == nil {
			lines := strings.Split(strings.ReplaceAll(string(output), "\r", ""), "\n")
			vals := make([]string, 0)
			skipHeader := true
			for _, line := range lines {
				trimmed := strings.TrimSpace(line)
				if trimmed == "" {
					skipHeader = false
					continue
				}
				if skipHeader {
					continue
				}
				if strings.HasPrefix(trimmed, "The command") || strings.HasPrefix(trimmed, "*") {
					continue
				}
				for _, member := range strings.Fields(line) {
					if len(member) > 0 {
						vals = append(vals, member)
					}
				}
			}
			return carapace.ActionValues(vals...)
		}
		return carapace.ActionValues()
	})
}
