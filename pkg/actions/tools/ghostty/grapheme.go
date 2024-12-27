package ghostty

import "github.com/carapace-sh/carapace"

// ActionGraphemeWidthMethods completes grapheme width methods
//
//	legacy (Use a legacy method to determine grapheme width)
//	unicode (Use the Unicode standard to determine grapheme width)
func ActionGraphemeWidthMethods() carapace.Action {
	return carapace.ActionValuesDescribed(
		"legacy", "Use a legacy method to determine grapheme width",
		"unicode", "Use the Unicode standard to determine grapheme width",
	)
}
