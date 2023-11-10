package ffmpeg

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionFilters() carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-filters")(func(output []byte) carapace.Action {
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
