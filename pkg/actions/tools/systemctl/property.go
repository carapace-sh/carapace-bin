package systemctl

import (
	"strings"

	"github.com/rsteube/carapace"
)

type PropertiesOpts struct {
	User  bool
	Units []string
}

// ActionProperties completes properties for given unit
//
//	Architecture (x86-64)
//	ConfirmSpawn (no)
func ActionProperties(opts PropertiesOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"show"}
		if opts.User {
			args = append(args, "user")
		}
		args = append(args, opts.Units...)
		return carapace.ActionExecCommand("systemctl", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				vals = append(vals, strings.SplitN(line, "=", 2)...)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
