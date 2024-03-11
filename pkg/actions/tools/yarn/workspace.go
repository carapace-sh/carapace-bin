package yarn

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionWorkspaces completes workspaces
//
//	example (.)
//	another (/tmp)
func ActionWorkspaces() carapace.Action {
	return actionYarn("workspaces", "list", "--json")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines[:len(lines)-1] {
			var workspace struct {
				Name     string
				Location string
			}
			if err := json.Unmarshal([]byte(line), &workspace); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			vals = append(vals, workspace.Name, workspace.Location)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
