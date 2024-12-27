package ghostty

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionShaderAnimationModes completes shader animation modes
//
//	true (run an animation loop when custom shaders are used)
//	false (only render when the terminal is updated)
func ActionShaderAnimationModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"true", "run an animation loop when custom shaders are used",
		"false", "only render when the terminal is updated",
		"always", "run the animation loop regardless of whether the terminal is focused or not",
	).StyleF(style.ForKeyword).Tag("shader animation modes")
}
