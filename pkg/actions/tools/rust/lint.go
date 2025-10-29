package rust

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionLintLevels completes lint levels
//
//	allow
//	deny
func ActionLintLevels() carapace.Action {
	return carapace.ActionStyledValues(
		"allow", style.Carapace.LogLevelInfo,
		"expect", style.Carapace.LogLevelDebug,
		"warn", style.Carapace.LogLevelWarning,
		"force-warn", style.Carapace.LogLevelError,
		"deny", style.Carapace.LogLevelCritical,
		"forbid", style.Carapace.LogLevelFatal,
	).Tag("lint levels").Uid("rust", "lint-level")
}
