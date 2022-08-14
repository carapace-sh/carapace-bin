// package task contains task related actions
package task

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionTasks completes tasks
//
//	default
//	build (build the project)
func ActionTasks(file string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"--list-all"}
		if file != "" {
			args = append(args, "--taskfile", file)
		}
		return carapace.ActionExecCommand("task", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^[*] (?P<name>[^ ]+):( \t*(?P<description>.*))?$`)

			vals := make([]string, 0)
			for _, line := range lines {
				if matches := r.FindStringSubmatch(line); matches != nil {
					vals = append(vals, matches[1], matches[3])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
