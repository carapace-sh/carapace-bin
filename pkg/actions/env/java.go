package env

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/conditions"
)

func init() {
	knownVariables["java"] = func() variables {
		return variables{
			Condition: conditions.ConditionPath("java"),
			Variables: map[string]string{
				"JAVA_HOME": "JDK installation directory",
			},
			VariableCompletion: map[string]carapace.Action{
				"JAVA_HOME": carapace.ActionDirectories(),
			},
		}
	}
}
