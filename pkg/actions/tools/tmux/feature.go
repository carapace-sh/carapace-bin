package tmux

import "github.com/rsteube/carapace"

// ActionFeatures completes features
//   `256 (Supports 256 colours with the SGR escape sequences.)
//   `RGB (Supports RGB colour with the SGR escape sequences.)
func ActionFeatures() carapace.Action {
	return carapace.ActionValuesDescribed(
		"256", "Supports 256 colours with the SGR escape sequences.",
		"clipboard", "Allows setting the system clipboard.",
		"ccolour", "Allows setting the cursor colour.",
		"cstyle", "Allows setting the cursor style.",
		"extkeys", "Supports extended keys.",
		"focus", "Supports focus reporting.",
		"margins", "Supports DECSLRM margins.",
		"mouse", "Supports xterm(1) mouse sequences.",
		"overline", "Supports the overline SGR attribute.",
		"rectfill", "Supports the DECFRA rectangle fill escape sequence.",
		"RGB", "Supports RGB colour with the SGR escape sequences.",
		"strikethrough", "Supports the strikethrough SGR escape sequence.",
		"sync", "Supports synchronized updates.",
		"title", "Supports xterm(1) title setting.",
		"usstyle", "Allows underscore style and colour to be set.",
	)
}
