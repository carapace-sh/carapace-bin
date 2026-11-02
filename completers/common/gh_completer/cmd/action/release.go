package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionReleases completes releases
func ActionReleases(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionReleases(gh.ReleaseOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}

// ActionNextReleases completes next release versions
func ActionNextReleases(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionNextReleases(gh.ReleaseOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}

// ActionReleaseAssets completes assets of a release
func ActionReleaseAssets(cmd *cobra.Command, tag string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionReleaseAssets(gh.ReleaseOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name, Tag: tag})
}
