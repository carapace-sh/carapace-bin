package glow

import "github.com/carapace-sh/carapace"

// ActionStyles completes styles
func ActionStyles() carapace.Action {
	return carapace.ActionValues(
		"ascii",
		"auto",
		"dark",
		"dracula",
		"light",
		"notty",
		"pink",
	).Tag("styles")
}
