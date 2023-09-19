package env

import "github.com/rsteube/carapace"

func init() {
	knownVariables["carapace"] = variables{
		Names: map[string]string{
			"CARAPACE_COVERDIR":      "coverage directory for sandbox tests",
			"CARAPACE_HIDDEN":        "show hidden commands/flags",
			"CARAPACE_LENIENT":       "allow unknown flags",
			"CARAPACE_LOG":           "enable logging",
			"CARAPACE_MATCH":         "match case insensitive",
			"CARAPACE_SANDBOX":       "mock context for sandbox tests",
			"CARAPACE_ZSH_HASH_DIRS": "zsh hash directories",
		},
		Values: map[string]carapace.Action{
			"CARAPACE_COVERDIR": carapace.ActionDirectories(),
			"CARAPACE_HIDDEN":   carapace.ActionValues("0", "1"),
			"CARAPACE_LENIENT":  carapace.ActionValues("0", "1"),
			"CARAPACE_LOG":      carapace.ActionValues("0", "1"),
			"CARAPACE_MATCH":    carapace.ActionValues("0", "1"),
		},
	}

}
