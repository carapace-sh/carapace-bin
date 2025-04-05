package action

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionColor() carapace.Action {
	return carapace.ActionExecCommand("gm", "convert", "-list", "color")(func(output []byte) carapace.Action {
		re := regexp.MustCompile(`^(\S+)\s+([0-9]+,\s*[0-9]+,\s*[0-9]+)\s+(.*)$`)

		lines := strings.Split(string(output), "\n")
		if len(lines) <= 4 {
			return carapace.ActionValues()
		}

		completions := make([]string, 0, len(lines)-4)

		for _, line := range lines[4:] {
			matches := re.FindStringSubmatch(strings.TrimSpace(line))
			if len(matches) != 4 {
				continue
			}

			value, rgb, compliance := matches[1], matches[2], strings.TrimSpace(matches[3])
			completions = append(completions, value, rgb+" ("+compliance+")")
		}

		return carapace.ActionValuesDescribed(completions...)
	})
}
