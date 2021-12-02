package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_forkCmd = &cobra.Command{
	Use:   "fork",
	Short: "Create a fork of a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_forkCmd).Standalone()
	repo_forkCmd.Flags().Bool("clone", false, "Clone the fork {true|false}")
	repo_forkCmd.Flags().String("org", "", "Create the fork in an organization")
	repo_forkCmd.Flags().Bool("remote", false, "Add remote for fork {true|false}")
	repo_forkCmd.Flags().String("remote-name", "origin", "Specify a name for a fork's new remote.")
	repoCmd.AddCommand(repo_forkCmd)

	carapace.Gen(repo_forkCmd).FlagCompletion(carapace.ActionMap{
		"org": action.ActionUsers(repo_forkCmd, action.UserOpts{Organizations: true}),
	})

	carapace.Gen(repo_forkCmd).PositionalCompletion(
		action.ActionOwnerRepositories(repo_forkCmd),
	)
}
