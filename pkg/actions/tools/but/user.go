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
	return carapace.ActionExecCommand("but", "config", "forge", "list-users")(func(output []byte) carapace.Action {
		r := regexp.MustCompile(`^\s*• (?:(?P<type>[^:]+): (?P<user>.+)|(?P<kind>Enterprise|Self-hosted) (?P<login>\S+@\S+)) \(.*\)$`)

		vals := make([]string, 0)
		for _, line := range strings.Split(string(output), "\n") {
			if matches := r.FindStringSubmatch(line); matches != nil {
				if user := matches[r.SubexpIndex("user")]; user != "" {
					vals = append(vals, user, matches[r.SubexpIndex("type")])
				} else {
					vals = append(vals, matches[r.SubexpIndex("login")], matches[r.SubexpIndex("kind")])
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
