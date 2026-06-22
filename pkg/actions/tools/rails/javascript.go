package rails

import "github.com/carapace-sh/carapace"

// ActionJavaScriptOptions completes rails new --javascript values
//
//	importmap
//	bun
func ActionJavaScriptOptions() carapace.Action {
	return carapace.ActionValues(
		"importmap",
		"bun",
		"webpack",
		"esbuild",
		"rollup",
	).Tag("javascript options").Uid("rails", "javascript-option")
}
