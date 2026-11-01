package action

import (
	"github.com/carapace-sh/carapace"
)

// ActionDeprecationLevels completes the deprecation warning levels for -deprecation
//
//	module:all
//	module:none
func ActionDeprecationLevels() carapace.Action {
	return carapace.ActionValues("module:all", "module:local", "module:none")
}

// ActionLintRules completes the lint rules for -lint
//
//	all
//	core:unused-variable
func ActionLintRules() carapace.Action {
	return carapace.ActionValues(
		"all",
		"core:all",
		"core:confusing",
		"core:count-instead-enabled",
		"core:improvement",
		"core:no-type-variable",
		"core:unused-local",
		"core:unused-variable",
	)
}
