package dagger

import "github.com/carapace-sh/carapace"

// ActionSdks completes skds
//
//	cue
//	dotnet
func ActionSdks() carapace.Action {
	return carapace.ActionValues(
		"cue",
		"dotnet",
		"elixir",
		"go",
		"java",
		"php",
		"python",
		"rust",
		"typescript",
	)
}
