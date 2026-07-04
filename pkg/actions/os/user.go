package os

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

func isNonLoginShell(shell string) bool {
	switch runtime.GOOS {
	case "darwin":
		return shell == "/usr/bin/false" || shell == "/usr/sbin/uucico"
	default:
		return shell == "/usr/bin/nologin" || shell == "/bin/false"
	}
}

// ActionUsers completes system user names
//
//	root (0)
//	daemon (1)
func ActionUsers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch runtime.GOOS {
		case "windows":
			return actionUsersWindows()
		default:
			return actionUsersUnix()
		}
	}).Tag("users").Uid("os", "user")
}

func actionUsersUnix() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		users := []string{}
		if content, err := os.ReadFile("/etc/passwd"); err == nil {
			for entry := range strings.SplitSeq(string(content), "\n") {
				splitted := strings.Split(entry, ":")
				if len(splitted) > 6 {
					user := splitted[0]
					id := splitted[2]
					description := strings.TrimSpace(splitted[4])
					shell := splitted[6]

					if description != "" {
						description = " - " + description
					}

					_style := style.Default
					if id == "0" {
						_style = style.Red
					} else if !isNonLoginShell(shell) {
						if _id, err := strconv.Atoi(id); err == nil && _id < 1000 {
							_style = style.Yellow
						} else {
							_style = style.Blue
						}
					}

					if len(strings.TrimSpace(user)) > 0 {
						users = append(users, user, id+description, _style)
					}
				}
			}
		}
		return carapace.ActionStyledValuesDescribed(users...)
	})
}

func actionUsersWindows() carapace.Action {
	return carapace.ActionExecCommand("net", "user")(func(output []byte) carapace.Action {
		lines := strings.Split(strings.ReplaceAll(string(output), "\r", ""), "\n")
		vals := make([]string, 0)
		// net user output has header lines, then user names in columns, then footer
		for i, line := range lines {
			if i < 4 { // skip header lines
				continue
			}
			trimmed := strings.TrimSpace(line)
			if trimmed == "" || strings.HasPrefix(trimmed, "The command") {
				continue
			}
			for _, user := range strings.Fields(line) {
				if len(user) > 0 {
					vals = append(vals, user, user, style.Blue)
				}
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}

// ActionUserGroup completes system user:group separately
//
//	bin:audio
//	lp:list
func ActionUserGroup() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionUsers().Invoke(c).Suffix(":").ToA()
		case 1:
			return ActionGroups()
		default:
			return carapace.ActionValues()
		}
	})
}
