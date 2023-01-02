package os

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionFontFamilies completes font family names
//
//	MesloLGSDZ Nerd Font
//	Nimbus Sans
func ActionFontFamilies() carapace.Action {
	return carapace.ActionExecCommand("fc-list", "--format=%{family}\n")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	}).Tag("font families")
}
