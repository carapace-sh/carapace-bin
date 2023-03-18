package yarn

import (
	"encoding/json"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionDependencies completes dependencies
//
//	project@workspace:.
//	yaml@npm:2.2.1
func ActionDependencies() carapace.Action {
	return carapace.ActionExecCommand("yarn", "info", "--json")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines[:len(lines)-1] {
			var dependency struct{ Value string }
			if err := json.Unmarshal([]byte(line), &dependency); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			vals = append(vals, dependency.Value)
		}
		return carapace.ActionValues(vals...)
	})
}
