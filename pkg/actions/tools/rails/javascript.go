package rails

import "github.com/carapace-sh/carapace"

// ActionJavaScriptOptions completes rails new --javascript values
//
//	importmap
//	bun
func ActionJavaScriptOptions() carapace.Action {
	return carapace.ActionValues(
		"bun",
		"esbuild",
		"importmap",
		"rollup",
		"webpack",
	).Tag("javascript options").Uid("rails", "javascript-option")
}
