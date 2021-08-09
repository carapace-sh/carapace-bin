package adb

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionDeviceUsers() carapace.Action {
	return carapace.ActionExecCommand("adb", "shell", "pm", "list", "users")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\tUserInfo{(?P<id>\d+):(?P<name>[^:]+):(?P<info>[^}]*)} (?P<state>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
