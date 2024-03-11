package yarn

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionPlugins completes plugins
//
//	@yarnpkg/plugin-constraints ([Yarn <4 only] Adds a new command to Yarn (`yarn constraints`) to enforce lin...)
//	@yarnpkg/plugin-exec ([Yarn <4 only] Adds a new protocol to Yarn (`exec:`) that dynamically generat...)
func ActionPlugins() carapace.Action {
	return carapace.ActionExecCommand("yarn", "plugin", "list", "--json")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines[:len(lines)-1] {
			var plugin struct {
				Name         string
				Experimental string
				Description  string
			}
			if err := json.Unmarshal([]byte(line), &plugin); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			if plugin.Experimental == "true" {
				vals = append(vals, plugin.Name, plugin.Description, style.Red)
			} else {
				vals = append(vals, plugin.Name, plugin.Description, style.Default)
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}
