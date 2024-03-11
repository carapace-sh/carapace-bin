package fastfetch

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionColors completes colors
//
//	black
//	red
func ActionColors() carapace.Action {
	return carapace.ActionStyledValues(
		"black", style.Black,
		"red", style.Red,
		"green", style.Green,
		"yellow", style.Yellow,
		"blue", style.Blue,
		"magenta", style.Magenta,
		"cyan", style.Cyan,
		"white", style.White,
		"default", style.Default,
	)
}
