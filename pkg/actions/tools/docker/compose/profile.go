package compose

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionProfiles completes profiles
//
//	debug
//	test
func ActionProfiles(files ...string) carapace.Action {
	return actionExecCompose(files, "config", "--profiles")(func(output []byte) carapace.Action {
		profiles := strings.Split(strings.TrimSpace(string(output)), "\n")
		if len(profiles) == 1 && profiles[0] == "" {
			return carapace.ActionValues()
		}
		return carapace.ActionValues(profiles...)
	})
}
