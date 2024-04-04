package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

func ActionVariables(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd, c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		if flag := cmd.Flag("org"); flag != nil && flag.Changed {
			return gh.ActionOrganisationVariables(gh.OwnerOpts{
				Host:  repo.RepoHost(),
				Owner: flag.Value.String(),
			})
		}

		if flag := cmd.Flag("env"); flag != nil && flag.Changed {
			return gh.ActionEnvironmentVariables(gh.EnvironmentOpts{
				Host:        repo.RepoHost(),
				Owner:       repo.RepoOwner(),
				Name:        repo.RepoName(),
				Environment: flag.Value.String(),
			})
		}

		return gh.ActionRepositoryVariables(gh.RepoOpts{
			Host:  repo.RepoHost(),
			Owner: repo.RepoOwner(),
			Name:  repo.RepoName(),
		})
	})
}
