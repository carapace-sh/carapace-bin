package scoop

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
)

type scoopAlias struct {
	Name    string `json:"Name"`
	Command string `json:"Command"`
}

// ActionAliases completes alias names with their command as description
//
//	rm (scoop uninstall $args[0])
//	upgrade (scoop update *)
func ActionAliases() carapace.Action {
	return carapace.ActionExecCommand("powershell", "-NoProfile", "-Command", "scoop alias list | ConvertTo-Json -Compress")(func(output []byte) carapace.Action {
		output = []byte(strings.TrimSpace(string(output)))
		if len(output) == 0 || string(output) == "null" {
			return carapace.ActionValues()
		}
		var aliases []scoopAlias
		if err := json.Unmarshal(output, &aliases); err != nil {
			var alias scoopAlias
			if err2 := json.Unmarshal(output, &alias); err2 != nil {
				return carapace.ActionMessage(err.Error())
			}
			return carapace.ActionValuesDescribed(alias.Name, alias.Command).Tag("aliases")
		}
		vals := make([]string, 0)
		for _, alias := range aliases {
			vals = append(vals, alias.Name, alias.Command)
		}
		return carapace.ActionValuesDescribed(vals...).Tag("aliases")
	})
}
