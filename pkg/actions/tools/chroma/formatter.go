package chroma

import "github.com/carapace-sh/carapace"

// ActionFormatters completes formatters
//
//	html
//	json
func ActionFormatters() carapace.Action {
	return carapace.ActionValues(
		"html",
		"json",
		"noop",
		"svg",
		"terminal",
		"terminal16",
		"terminal16m",
		"terminal256",
		"terminal8",
		"tokens",
	).Tag("formatters")
}
