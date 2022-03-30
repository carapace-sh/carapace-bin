package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

// ActionSearchMultiRepo handles the stringslice `repo` flag of the search commands.
// For each repo given function is invoked with a fake repo flag and the return values are merged.
func ActionSearchMultiRepo(cmd *cobra.Command, f func(cmd *cobra.Command) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repos, err := cmd.Flags().GetStringSlice("repo")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		if len(repos) == 0 {
			repos = append(repos, "")
		}

		batch := carapace.Batch()
		for _, repo := range repos {
			name := repo
			batch = append(batch, carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				dummyCmd := &cobra.Command{}
				dummyCmd.Flags().String("repo", name, "fake repo flag")
				if name != "" {
					dummyCmd.Flag("repo").Changed = true
				}
				return f(dummyCmd)
			}))
		}
		return batch.ToA()
	})
}

func ActionSearchIssueFields() carapace.Action {
	return carapace.ActionValues(
		"assignees",
		"author",
		"authorAssociation",
		"body",
		"closedAt",
		"commentsCount",
		"createdAt",
		"id",
		"isLocked",
		"isPullRequest",
		"labels",
		"number",
		"repository",
		"state",
		"title",
		"updatedAt",
		"url",
	)
}

func ActionSearchRepositoryFields() carapace.Action {
	return carapace.ActionValues(
		"createdAt",
		"defaultBranch",
		"description",
		"forksCount",
		"fullName",
		"hasDownloads",
		"hasIssues",
		"hasPages",
		"hasProjects",
		"hasWiki",
		"homepage",
		"id",
		"isArchived",
		"isDisabled",
		"isFork",
		"isPrivate",
		"language",
		"license",
		"name",
		"openIssuesCount",
		"owner",
		"pushedAt",
		"size",
		"stargazersCount",
		"updatedAt",
		"visibility",
		"watchersCount",
	)
}
