package ffmpeg

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionEncoders completes encoders
//
//	ac3 (ATSC A/52A (AC-3))
//	ac3_fixed (ATSC A/52A (AC-3) (codec ac3))
func ActionEncoders() carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-encoders")(func(output []byte) carapace.Action {
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
	}).Tag("encoders")
}
