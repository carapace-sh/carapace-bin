// package mix contains mix related actions
package mix

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionMixTasks completes mix tasks
//
//	install
//	deps.get
func ActionMixTasks() carapace.Action {
	return carapace.ActionExecCommand("mix", "help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^mix (\S+)\s*# (.*)`)
		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("tasks")
}
