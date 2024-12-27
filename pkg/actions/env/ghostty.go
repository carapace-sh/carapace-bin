package env

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/conditions"
)

func init() {
	knownVariables["ghostty"] = func() variables {
		return variables{
			Condition: conditions.ConditionPath("ghostty"),
			Variables: map[string]string{
				"GHOSTTY_RESOURCES_DIR": "resources directory",
			},
			VariableCompletion: map[string]carapace.Action{
				"GHOSTTY_RESOURCES_DIR": carapace.ActionDirectories(),
			},
		}
	}
}
