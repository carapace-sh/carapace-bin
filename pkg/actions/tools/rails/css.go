package rails

import "github.com/carapace-sh/carapace"

// ActionCSSOptions completes rails new --css values
//
//	tailwind
//	bootstrap
func ActionCSSOptions() carapace.Action {
	return carapace.ActionValues(
		"tailwind",
		"bootstrap",
		"bulma",
		"postcss",
		"sass",
	).Tag("css options").Uid("rails", "css-option")
}
