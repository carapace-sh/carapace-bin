package but

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionUsers completes users
//
//	userA (OAuth)
//	userB (OAuth)
func ActionUsers() carapace.Action {
	return carapace.ActionExecCommand("but", "forge", "list-users")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^- (?P<type>[^:]+): (?P<user>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[2], matches[1])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
