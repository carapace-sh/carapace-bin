package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionMentionableUsers completes mentionable users
func ActionMentionableUsers(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionMentionableUsers(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}

// ActionAssignableUsers completes assignable users
func ActionAssignableUsers(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionAssignableUsers(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}
