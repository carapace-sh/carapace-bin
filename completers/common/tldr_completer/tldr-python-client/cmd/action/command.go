package action

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionCommands(language string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if language == "" {
			language = "en"
		}
		return carapace.ActionExecCommand("tldr", "--list", "--language", language)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<command>.*) \((?P<language>.*)\)$`)

			vals := make([]string, 0)
			for _, line := range lines {
				if matches := r.FindStringSubmatch(line); matches != nil {
					vals = append(vals, matches[1], matches[2])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
