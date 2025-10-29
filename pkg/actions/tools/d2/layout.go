package d2

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionLayouts completes layouts
//
//	dagre (The directed graph layout library Dagre)
//	elk (Eclipse Layout Kernel (ELK) with the Layered algorithm.)
func ActionLayouts() carapace.Action {
	return carapace.ActionExecCommand("d2", "layout")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<name>[^ ]+) (\(bundled\) )?- (?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if line == "Usage:" {
				break
			}
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[3])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("layouts")
}
