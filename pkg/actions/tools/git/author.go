package git

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionAuthors completes authors of repo
//
//	Some Name (email@host.com)
//	Another Name (email@another.com)
func ActionAuthors() carapace.Action {
	return carapace.ActionExecCommand("git", "shortlog", "--summary", "--email", "HEAD")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^ +\d+\t(?P<name>.*) <(?P<email>[^<>]+)>$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				vals = append(vals, r.FindStringSubmatch(line)[1:]...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
