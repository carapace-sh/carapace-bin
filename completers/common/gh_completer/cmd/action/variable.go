package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionVariables completes variables (repository, environment or organization)
func ActionVariables(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	repoHost := repo.Host
	repoOwner := repo.Owner
	repoName := repo.Name

	if flag := cmd.Flag("org"); flag != nil && flag.Changed {
		return gh.ActionOrganizationVariables(gh.OwnerOpts{
			Host:  repoHost,
			Owner: flag.Value.String(),
		})
	}

	if flag := cmd.Flag("env"); flag != nil && flag.Changed {
		return gh.ActionEnvironmentVariables(gh.EnvironmentOpts{
			Host:        repoHost,
			Owner:       repoOwner,
			Name:        repoName,
			Environment: flag.Value.String(),
		})
	}

	return gh.ActionRepositoryVariables(gh.RepoOpts{
		Host:  repoHost,
		Owner: repoOwner,
		Name:  repoName,
	})
}
