package tofu

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionResources completes resources
//
//	aws_instance.example
//	aws_s3_bucket.example
func ActionResources(state string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"state", "list", "-no-color"}
		if state != "" {
			args = append(args, "-state", state)
		}
		return carapace.ActionExecCommand("tofu", args...)(func(output []byte) carapace.Action {
			vals := make([]string, 0)
			for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
				if line != "" {
					vals = append(vals, line)
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}
