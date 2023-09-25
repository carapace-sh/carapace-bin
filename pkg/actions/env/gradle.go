package env

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/conditions"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
)

func init() {
	knownVariables["gradle"] = variables{
		Condition: conditions.ConditionPath("gradle"),
		Variables: map[string]string{
			"GRADLE_HOME":      "Gradle distribution location",
			"GRADLE_OPTS":      "JVM arguments to use when starting the Gradle client VM",
			"GRADLE_USER_HOME": "Gradle User Home directory (",
		},
		VariableCompletion: map[string]carapace.Action{
			"GRADLE_HOME":      carapace.ActionDirectories(),
			"GRADLE_OPTS":      bridge.ActionCarapaceBin("java").Split(),
			"GRADLE_USER_HOME": carapace.ActionDirectories(),
		},
	}

}
