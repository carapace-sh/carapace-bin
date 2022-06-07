package git

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionCommitters completes committers of repo
//   GitHub
//   another
func ActionCommitters() carapace.Action {
	return carapace.ActionExecCommand("git", "shortlog", "--summary", "--committer", "HEAD")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^ +\d+\t(?P<name>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				vals = append(vals, r.FindStringSubmatch(line)[1])
			}
		}
		return carapace.ActionValues(vals...)
	})
}
