package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_deleteCmd = &cobra.Command{
	Use:     "delete [<repository>]",
	Short:   "Delete a repository",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_deleteCmd).Standalone()

	repo_deleteCmd.Flags().Bool("confirm", false, "Confirm deletion without prompting")
	repo_deleteCmd.Flags().Bool("yes", false, "Confirm deletion without prompting")
	repo_deleteCmd.Flag("confirm").Hidden = true
	repoCmd.AddCommand(repo_deleteCmd)

	carapace.Gen(repo_deleteCmd).PositionalCompletion(
		action.ActionOwnerRepositories(repo_deleteCmd),
	)
}
