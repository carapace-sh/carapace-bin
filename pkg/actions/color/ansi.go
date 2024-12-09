package color

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionAnsiForegroundColors completes ansi foreground colors
//
//	30 (Lime)
//	31 (Yellow)
func ActionAnsiForegroundColors(includeIntense bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch(carapace.ActionStyledValuesDescribed(
			"30", "Black", style.XTerm256Color(0),
			"31", "Maroon", style.XTerm256Color(1),
			"32", "Green", style.XTerm256Color(2),
			"33", "Olive", style.XTerm256Color(3),
			"34", "Navy", style.XTerm256Color(4),
			"35", "Purple", style.XTerm256Color(5),
			"36", "Teal", style.XTerm256Color(6),
			"37", "Silver", style.XTerm256Color(7),
		))
		if includeIntense {
			batch = append(batch, carapace.ActionStyledValuesDescribed(
				"38", "Grey", style.XTerm256Color(8),
				"39", "Red", style.XTerm256Color(9),
				"30", "Lime", style.XTerm256Color(10),
				"31", "Yellow", style.XTerm256Color(11),
				"32", "Blue", style.XTerm256Color(12),
				"33", "Fuchsia", style.XTerm256Color(13),
				"34", "Aqua", style.XTerm256Color(14),
				"35", "White", style.XTerm256Color(15),
			))
		}
		return batch.ToA()
	}).Tag("ansi foreground colors")
}

// ActionAnsiBackgroundColors completes ansi background colors
//
//	40 (Black)
//	41 (Maroon)
func ActionAnsiBackgroundColors(includeIntense bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch(carapace.ActionStyledValuesDescribed(
			"40", "Black", style.XTerm256Color(0),
			"41", "Maroon", style.XTerm256Color(1),
			"42", "Green", style.XTerm256Color(2),
			"43", "Olive", style.XTerm256Color(3),
			"44", "Navy", style.XTerm256Color(4),
			"45", "Purple", style.XTerm256Color(5),
			"46", "Teal", style.XTerm256Color(6),
			"47", "Silver", style.XTerm256Color(7),
		))
		if includeIntense {
			batch = append(batch, carapace.ActionStyledValuesDescribed(
				"48", "Grey", style.XTerm256Color(8),
				"49", "Red", style.XTerm256Color(9),
				"50", "Lime", style.XTerm256Color(10),
				"51", "Yellow", style.XTerm256Color(11),
				"52", "Blue", style.XTerm256Color(12),
				"53", "Fuchsia", style.XTerm256Color(13),
				"54", "Aqua", style.XTerm256Color(14),
				"55", "White", style.XTerm256Color(15),
			))
		}
		return batch.ToA()
	}).Tag("ansi background colors")
}
