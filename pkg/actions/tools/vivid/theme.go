package vivid

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionThemes completes themes
//
//	alabaster_dark
//	ayu
func ActionThemes() carapace.Action {
	return carapace.ActionExecCommand("vivid", "themes")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
