package os

import (
	"os"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionUsers completes system user names
//   root (0)
//   daemon (1)
func ActionUsers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		users := []string{}
		if content, err := os.ReadFile("/etc/passwd"); err == nil {
			for _, entry := range strings.Split(string(content), "\n") {
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
					} else if shell != "/usr/bin/nologin" {
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

// ActionUserGroup completes system user:group separately
//   bin:audio
//   lp:list
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
