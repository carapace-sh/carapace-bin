package dagger

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
)

// ActionMods copletes local files and remote repositories
func ActionMods() carapace.Action {
	return carapace.Batch(
		carapace.ActionFiles(),
		git.ActionRepositorySearch(git.SearchOpts{Github: true}),
	).ToA()
}
