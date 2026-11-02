package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionCaches completes caches
func ActionCaches(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionCaches(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}

// ActionCacheRefs completes cache refs
func ActionCacheRefs(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionCacheRefs(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}
