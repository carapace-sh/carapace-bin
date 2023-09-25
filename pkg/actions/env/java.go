package env

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/conditions"
)

func init() {
	knownVariables["java"] = variables{
		Condition: conditions.ConditionPath("java"),
		Variables: map[string]string{
			"JAVA_HOME": "JDK installation directory",
		},
		VariableCompletion: map[string]carapace.Action{
			"JAVA_HOME": carapace.ActionDirectories(),
		},
	}

}
