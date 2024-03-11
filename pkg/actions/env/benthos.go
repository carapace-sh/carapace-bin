package env

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/conditions"
)

func init() {
	knownVariables["benthos"] = func() variables {
		return variables{
			Condition: conditions.ConditionPath("benthos"),
			Variables: map[string]string{
				"BSTDIO_NODE_TOKEN":  "A token for the session, used to authenticate requests",
				"BSTDIO_NODE_SECRET": "A token secret the session, used to authenticate requests",
			},
			VariableCompletion: map[string]carapace.Action{
				"BSTDIO_NODE_TOKEN":  carapace.ActionValues(),
				"BSTDIO_NODE_SECRET": carapace.ActionValues(),
			},
		}
	}
}
