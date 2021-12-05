package gh

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

// ActionOwnerRepositories completes owner/repository from github.com separately
//   cli (GitHub CLI)
//   cli/cli ((GitHub’s official command line tool))
func ActionOwnerRepositories() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return action.ActionOwnerRepositories(&cobra.Command{})
	})
}

func ActionUsers(users bool, orgs bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return action.ActionUsers(&cobra.Command{}, action.UserOpts{Users: users, Organizations: orgs})
	})
}
