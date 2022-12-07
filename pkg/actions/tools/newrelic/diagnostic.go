package newrelic

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionDiagnosticSuites completes diagnostic suites
//
//	dotnetcore (.NET Core Agent installation)
//	infra (Infrastructure Agent installation)
func ActionDiagnosticSuites() carapace.Action {
	return carapace.ActionExecCommand("newrelic", "diagnose", "run", "--list-suites")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<name>[^ :]+)    +(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
