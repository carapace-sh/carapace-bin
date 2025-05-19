package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_forkCmd = &cobra.Command{
	Use:   "fork <repo>",
	Short: "Fork a GitLab repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_forkCmd).Standalone()

	repo_forkCmd.Flags().BoolP("clone", "c", false, "Clone the fork. Options: true, false.")
	repo_forkCmd.Flags().StringP("name", "n", "", "The name assigned to the new project after forking.")
	repo_forkCmd.Flags().StringP("path", "p", "", "The path assigned to the new project after forking.")
	repo_forkCmd.Flags().Bool("remote", false, "Add a remote for the fork. Options: true, false.")
	repoCmd.AddCommand(repo_forkCmd)

	// TODO path completion

	carapace.Gen(repo_forkCmd).PositionalCompletion(
		action.ActionRepo(repo_forkCmd),
	)
}
