package os

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
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
				if len(splitted) > 2 {
					user := splitted[0]
					id := splitted[2]
					if len(strings.TrimSpace(user)) > 0 {
						users = append(users, user, id)
					}
				}
			}
		}
		return carapace.ActionValuesDescribed(users...)
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
