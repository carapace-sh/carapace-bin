package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_deleteCmd = &cobra.Command{
	Use:   "delete <tag>",
	Short: "Delete a GitLab release.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_deleteCmd).Standalone()

	release_deleteCmd.Flags().BoolP("with-tag", "t", false, "Delete the associated tag.")
	release_deleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt.")
	releaseCmd.AddCommand(release_deleteCmd)

	carapace.Gen(release_deleteCmd).PositionalCompletion(
		action.ActionReleases(release_deleteCmd),
	)
}
