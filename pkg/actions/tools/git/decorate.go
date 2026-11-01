package git

import "github.com/carapace-sh/carapace"

// ActionDecorateModes completes decorate modes
//
//	short (do not print ref prefixes)
//	full (print ref prefixes)
func ActionDecorateModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"auto", "short format when output to terminal",
		"full", "print ref prefixes",
		"no", "no decoration",
		"short", "do not print ref prefixes",
	).Tag("decorate modes").Uid("git", "decorate-mode")
}
