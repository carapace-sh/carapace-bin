package git

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
)

// ActionHooks completes hooks
//
//	pre-commit.sample
//	pre-merge-commit.sample
func ActionHooks() carapace.Action {
	return carapace.ActionExecutables("hooks").
		ChdirF(traverse.GitDir).
		Tag("hooks")
}
