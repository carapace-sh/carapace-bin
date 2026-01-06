package ghostty

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionFontSyntheticStyles completes font synthetic styles
//
//	false (completely disable)
//	italic (apply a slant to the glyph)
func ActionFontSyntheticStyles() carapace.Action {
	return carapace.ActionValuesDescribed(
		"bold", "draw an outline around the glyph",
		"bold-italic", "draw an outline and apply a slant to the glyph",
		"false", "completely disable",
		"italic", "apply a slant to the glyph",
		"no-bold", "disable for bold",
		"no-bold-italic", "disable for bold-italic",
		"no-italic", "disable for italic",
		"true", "completely enable",
	).Tag("font synthetic styles")
}

// ActionFreetypeLoadFlags completes freetype load flags
//
//	hinting (Enable hinting)
//	force-autohint (Use the freetype auto-hinter)
func ActionFreetypeLoadFlags() carapace.Action {
	return carapace.ActionValuesDescribed(
		"hinting", "Enable hinting",
		"force-autohint", "Use the freetype auto-hinter",
		"monochrome", "use 1-bit monochrome rendering",
		"autohint", "Use the freetype auto-hinter",
		"no-hinting", "Disable hinting",
		"no-force-autohint", "Do not se the freetype auto-hinter",
		"no-monochrome", "Do not use 1-bit monochrome rendering",
		"no-autohint", "Do not use the freetype auto-hinter",
	).Tag("freetype load flags")
}

// ActionFontFamilies completes font families
//
//	FreeMono
//	FreeSans
func ActionFontFamilies() carapace.Action {
	return carapace.ActionExecCommand("ghostty", "+list-fonts")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if line != "" && !strings.HasPrefix(line, " ") {
				vals = append(vals, line)
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("font families")
}
