package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_archiveCmd = &cobra.Command{
	Use:     "archive [<repository>]",
	Short:   "Archive a repository",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_archiveCmd).Standalone()

	repo_archiveCmd.Flags().Bool("confirm", false, "Skip the confirmation prompt")
	repo_archiveCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	repo_archiveCmd.Flag("confirm").Hidden = true
	repoCmd.AddCommand(repo_archiveCmd)

	carapace.Gen(repo_archiveCmd).PositionalCompletion(
		action.ActionOwnerRepositories(repo_archiveCmd),
	)
}
