package chdman

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionTemplates completes templates
//
//	0 (Conner - CFA170A)
//	1 (Rodime - R0201)
func ActionTemplates() carapace.Action {
	return carapace.ActionExecCommand("chdman", "listtemplates")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[4:] {
			if fields := strings.Fields(line); len(fields) > 3 {
				vals = append(vals, fields[0], fmt.Sprintf("%v - %v", fields[1], fields[2]))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("templates")
}
