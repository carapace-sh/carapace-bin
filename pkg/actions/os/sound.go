package os

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionSoundCards completes sound cards
//
//	0 (HDMI)
//	PCH (1)
func ActionSoundCards() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("aplay", "-l")(func(output []byte) carapace.Action {
			cards := make(map[string]string)
			r := regexp.MustCompile(`^card (?P<id>\d+): (?P<name>.+) \[.*\], device (?P<device>\d+).*$`)
			for _, line := range strings.Split(string(output), "\n") {
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
	}).Tag("soundcards")
}
