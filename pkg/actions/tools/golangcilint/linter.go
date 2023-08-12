package golangcilint

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionLinters completes linters
//
//	deadcode (Finds unused code)
//	dogsled (Checks assignments with too many blank identifiers)
func ActionLinters() carapace.Action {
	return carapace.ActionExecCommand("golangci-lint", "linters")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<linter>[^: (\[]+)[^:]*: (?P<description>[^\[(]+)`)

		s := style.Green
		vals := make([]string, 0)
		for _, line := range lines {
			if strings.HasPrefix(line, "Disabled") {
				s = style.Red
			}
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2], s)
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}
