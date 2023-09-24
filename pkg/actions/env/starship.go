package env

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/conditions"
	"github.com/rsteube/carapace/pkg/style"
)

func init() {
	knownVariables["starship"] = variables{
		Condition: conditions.ConditionPath("starship"),
		Variables: map[string]string{
			"STARSHIP_CACHE":       "cache location",
			"STARSHIP_CONFIG":      "config location",
			"STARSHIP_LOG":         "log level",
			"STARSHIP_NUM_THREADS": "number of threads",
			"STARSHIP_SESSION_KEY": "session key",
			"STARSHIP_SHELL":       "shell",
		},
		VariableCompletion: map[string]carapace.Action{
			"STARSHIP_CACHE":  carapace.ActionDirectories(),
			"STARSHIP_CONFIG": carapace.ActionFiles(),
			"STARSHIP_LOG":    carapace.ActionValues("debug", "error", "info", "trace", "warn").StyleF(style.ForLogLevel),
			"STARSHIP_SHELL": carapace.ActionValues(
				"bash",
				"cmd",
				"elvish",
				"fish",
				"ion",
				"nu",
				"powershell",
				"pwsh",
				"tcsh",
				"xonsh",
				"zsh",
			),
		},
	}

}
