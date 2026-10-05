package tofu

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionOutputs completes output values
//
//	instance_ip
//	db_password
func ActionOutputs(state string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"output", "-no-color"}
		if state != "" {
			args = append(args, "--state", state)
		}
		return carapace.ActionExecCommand("tofu", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<name>[^ =]+) = `)

			vals := make([]string, 0)
			for _, line := range lines {
				if r.MatchString(line) {
					vals = append(vals, r.FindStringSubmatch(line)[1])
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}
