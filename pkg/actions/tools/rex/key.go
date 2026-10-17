package rex

import (
	"github.com/carapace-sh/carapace"
)

// ActionKeys completes key names for `rex send-key`
//
//	enter
//	ctrl+b
func ActionKeys() carapace.Action {
	return carapace.ActionMultiParts("+", func(c carapace.Context) carapace.Action {
		baseKeys := []string{
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
			"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
			"backspace", "delete", "down", "end", "enter", "escape", "home",
			"insert", "left", "pagedown", "pageup", "right", "space", "tab", "up",
		}
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				carapace.ActionValues("alt", "cmd", "ctrl", "shift").Suffix("+"),
				carapace.ActionValues(baseKeys...),
			).ToA()
		default:
			return carapace.ActionValues(baseKeys...)
		}
	}).Tag("keys")
}
