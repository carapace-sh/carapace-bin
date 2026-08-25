package git

import "github.com/carapace-sh/carapace"

// ActionDecorateModes completes decorate modes
//
//	short (do not print ref prefixes)
//	full (print ref prefixes)
func ActionDecorateModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"short", "do not print ref prefixes",
		"full", "print ref prefixes",
		"auto", "short format when output to terminal",
		"no", "no decoration",
	).Tag("decorate modes").Uid("git", "decorate-mode")
}
