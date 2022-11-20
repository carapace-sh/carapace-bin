package golang

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
)

// ActionModuleSearch completes git repos in module format
//
//	github.com/rsteube/carapace@v0.0.1
//	github.com/spf13/cobra@v0.0.2
func ActionModuleSearch() carapace.Action {
	return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			opts := git.SearchOpts{}.Default()
			opts.Prefix = false
			return git.ActionRepositorySearch(opts).NoSpace()
		case 1:
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: "https://" + c.Parts[0], Tags: true})
		default:
			return carapace.ActionValues()
		}
	})
}
