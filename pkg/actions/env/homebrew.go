package env

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/conditions"
)

func init() {
	knownVariables["homebrew"] = func() variables {
		return variables{
			Condition: conditions.ConditionPath("brew"),
			Variables: map[string]string{
				"HOMEBREW_EVAL_ALL": "Evaluate all available formulae and casks",
			},
			VariableCompletion: map[string]carapace.Action{
				"HOMEBREW_EVAL_ALL": carapace.ActionValues("1"),
			},
		}
	}
}
