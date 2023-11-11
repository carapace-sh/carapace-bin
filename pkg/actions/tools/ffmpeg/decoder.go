package ffmpeg

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionDecoders completes decoders
//
//	4xm (4X Movie)
//	8bps (QuickTime 8BPS video)
func ActionDecoders() carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-decoders")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^ (?P<type>.).{5} (?P<name>[^ ]+) +(?P<description>.*)$`)

		found := false
		vals := make([]string, 0)
		for _, line := range lines {
			if !found {
				found = line == " ------"
				continue
			}

			if matches := r.FindStringSubmatch(line); matches != nil {
				switch matches[1] {
				case "V":
					vals = append(vals, matches[2], matches[3], style.Blue)
				case "A":
					vals = append(vals, matches[2], matches[3], style.Yellow)
				case "S":
					vals = append(vals, matches[2], matches[3], style.Magenta)
				}
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("decoders")
}
