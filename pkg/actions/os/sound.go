package os

import (
	"regexp"
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/condition"
)

// ActionSoundCards completes sound cards
//
//	0 (HDMI)
//	PCH (1)
func ActionSoundCards() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch {
		case condition.Executable("aplay")(c):
			return carapace.ActionExecCommand("aplay", "-l")(func(output []byte) carapace.Action {
				cards := make(map[string]string)
				r := regexp.MustCompile(`^card (?P<id>\d+): (?P<name>.+) \[.*\], device (?P<device>\d+).*$`)
				for line := range strings.SplitSeq(string(output), "\n") {
					if r.MatchString(line) {
						matches := r.FindStringSubmatch(line)
						cards[matches[1]] = matches[2]
					}
				}

				vals := make([]string, 0, len(cards)*2)
				for id, name := range cards {
					vals = append(vals, id, name)
					vals = append(vals, name, id)
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		case runtime.GOOS == "darwin":
			return carapace.ActionExecCommand("system_profiler", "SPAudioDataType", "-detailLevel", "mini")(func(output []byte) carapace.Action {
				vals := make([]string, 0)
				for line := range strings.SplitSeq(string(output), "\n") {
					line = strings.TrimSpace(line)
					if prefix, value, ok := strings.Cut(line, ":"); ok && strings.TrimSpace(prefix) == "device" {
						name := strings.TrimSpace(value)
						if name != "" {
							vals = append(vals, name, name)
						}
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		default:
			return carapace.ActionValues() // Windows and other platforms: no standard sound card enumeration
		}
	}).Tag("soundcards")
}
