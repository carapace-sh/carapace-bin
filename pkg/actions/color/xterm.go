package color

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionXtermColorNames completes xterm color names
//
//	Green
//	Olive
func ActionXtermColorNames() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionStyledValues(
			"Black", style.XTerm256Color(0),
			"Maroon", style.XTerm256Color(1),
			"Green", style.XTerm256Color(2),
			"Olive", style.XTerm256Color(3),
			"Navy", style.XTerm256Color(4),
			"Purple", style.XTerm256Color(5),
			"Teal", style.XTerm256Color(6),
			"Silver", style.XTerm256Color(7),
			"Grey", style.XTerm256Color(8),
			"Red", style.XTerm256Color(9),
			"Lime", style.XTerm256Color(10),
			"Yellow", style.XTerm256Color(11),
			"Blue", style.XTerm256Color(12),
			"Fuchsia", style.XTerm256Color(13),
			"Aqua", style.XTerm256Color(14),
			"White", style.XTerm256Color(15),
		).Tag("xterm colors")
	})
}
