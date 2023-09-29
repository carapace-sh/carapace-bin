package env

import (
	"github.com/rsteube/carapace"
)

func init() {
	knownVariables["xdg"] = func() variables {
		return variables{
			Variables: map[string]string{
				"XDG_DATA_HOME":   "base directory relative to which user-specific data files should be stored",
				"XDG_CONFIG_HOME": "base directory relative to which user-specific configuration files should be stored",
				"XDG_STATE_HOME":  "base directory relative to which user-specific state files should be stored",
				"XDG_DATA_DIRS":   "base directories to search for data files",
				"XDG_CONFIG_DIRS": "base directories to search for configuration files",
				"XDG_CACHE_HOME":  "base directory relative to which user-specific non-essential data files should be stored",
				"XDG_RUNTIME_DIR": "base directory relative to which user-specific non-essential runtime files should be stored",
			},
			VariableCompletion: map[string]carapace.Action{
				"XDG_DATA_HOME":   carapace.ActionDirectories(),
				"XDG_CONFIG_HOME": carapace.ActionDirectories(),
				"XDG_STATE_HOME":  carapace.ActionDirectories(),
				"XDG_DATA_DIRS":   carapace.ActionDirectories().List(":"),
				"XDG_CONFIG_DIRS": carapace.ActionDirectories().List(":"),
				"XDG_CACHE_HOME":  carapace.ActionDirectories(),
				"XDG_RUNTIME_DIR": carapace.ActionDirectories(),
			},
		}
	}
}
