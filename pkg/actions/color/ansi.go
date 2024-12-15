package color

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionAnsiForegroundColors completes ansi foreground colors
//
//	30 (Black)
//	31 (Red)
func ActionAnsiForegroundColors(includeIntense bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch(carapace.ActionStyledValuesDescribed(
			"30", "Black", style.XTerm256Color(0),
			"31", "Red", style.XTerm256Color(1),
			"32", "Green", style.XTerm256Color(2),
			"33", "Yellow", style.XTerm256Color(3),
			"34", "Blue", style.XTerm256Color(4),
			"35", "Magenta", style.XTerm256Color(5),
			"36", "Cyan", style.XTerm256Color(6),
			"37", "White", style.XTerm256Color(7),
		))
		if includeIntense {
			batch = append(batch, carapace.ActionStyledValuesDescribed(
				"90", "Bright Black", style.XTerm256Color(8),
				"91", "Bright Red", style.XTerm256Color(9),
				"92", "Bright Green", style.XTerm256Color(10),
				"93", "Bright Yellow", style.XTerm256Color(11),
				"94", "Bright Blue", style.XTerm256Color(12),
				"95", "Bright Magenta", style.XTerm256Color(13),
				"96", "Bright Cyan", style.XTerm256Color(14),
				"97", "Bright White", style.XTerm256Color(15),
			))
		}
		return batch.ToA()
	}).Tag("ansi foreground colors")
}

// ActionAnsiBackgroundColors completes ansi background colors
//
//	40 (Black)
//	41 (Red)
func ActionAnsiBackgroundColors(includeIntense bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch(carapace.ActionStyledValuesDescribed(
			"40", "Black", style.XTerm256Color(0),
			"41", "Red", style.XTerm256Color(1),
			"42", "Green", style.XTerm256Color(2),
			"43", "Yellow", style.XTerm256Color(3),
			"44", "Blue", style.XTerm256Color(4),
			"45", "Magenta", style.XTerm256Color(5),
			"46", "Cyan", style.XTerm256Color(6),
			"47", "White", style.XTerm256Color(7),
		))
		if includeIntense {
			batch = append(batch, carapace.ActionStyledValuesDescribed(
				"100", "Bright Black", style.XTerm256Color(8),
				"101", "Bright Red", style.XTerm256Color(9),
				"102", "Bright Green", style.XTerm256Color(10),
				"103", "Bright Yellow", style.XTerm256Color(11),
				"104", "Bright Blue", style.XTerm256Color(12),
				"105", "Bright Magenta", style.XTerm256Color(13),
				"106", "Bright Cyan", style.XTerm256Color(14),
				"107", "Bright White", style.XTerm256Color(15),
			))
		}
		return batch.ToA()
	}).Tag("ansi background colors")
}
