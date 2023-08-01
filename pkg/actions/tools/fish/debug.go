package fish

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionDebugCategories completes debug categories
//
//	reader-render (Rendering the command line)
//	screen (Screen repaints)
func ActionDebugCategories() carapace.Action {
	return carapace.ActionExecCommand("fish", "--print-debug-categories")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 1 {
				vals = append(vals, fields[0], strings.Join(fields[1:], " "))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
