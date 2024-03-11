package ffmpeg

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionFilters completes filters
//
//	acrusher (Reduce audio bit resolution.)
//	acue (Delay filtering to match a cue.)
func ActionFilters() carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-filters")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^ .{3} (?P<name>[^ ]+) +[^ ]+ *(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("filters")
}
