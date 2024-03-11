package env

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/conditions"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
)

func init() {
	knownVariables["jj"] = func() variables {
		return variables{
			Condition: conditions.ConditionPath("jj"),
			Variables: map[string]string{
				"JJ_CONFIG": "The location of the jj config file",
				"JJ_EDITOR": "The default editor",
			},
			VariableCompletion: map[string]carapace.Action{
				"JJ_CONFIG": carapace.ActionFiles(),
				"JJ_EDITOR": bridge.ActionCarapaceBin().Split(),
			},
		}
	}
}
