package action

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionColumns() carapace.Action {
	return carapace.ActionExecCommand("lsfd", "--list-columns")(func(output []byte) carapace.Action {
		re := regexp.MustCompile(`^\s*(\S+)\s+<[^>]+>\s+(.*)$`)
		var values []string

		for _, line := range strings.Split(string(output), "\n") {
			if matches := re.FindStringSubmatch(line); len(matches) == 3 {
				name, description := matches[1], matches[2]
				values = append(values, name, description)
			}
		}

		return carapace.ActionValuesDescribed(values...)
	})
}
