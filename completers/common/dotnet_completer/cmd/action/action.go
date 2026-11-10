package action

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

var columnSeparator = regexp.MustCompile(`\s{2,}`)

// ActionTemplates completes template short names for `dotnet new`
//
//	console
//	webapi
func ActionTemplates() carapace.Action {
	return carapace.ActionExecCommand("dotnet", "new", "list")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		inTable := false
		for _, line := range strings.Split(string(output), "\n") {
			if !inTable {
				inTable = strings.HasPrefix(line, "---")
				continue
			}
			if strings.TrimSpace(line) == "" {
				break
			}

			columns := columnSeparator.Split(strings.TrimSpace(line), -1)
			if len(columns) < 2 {
				continue
			}
			for _, shortName := range strings.Split(columns[1], ",") {
				vals = append(vals, strings.TrimSpace(shortName), columns[0])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
