package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_deleteCmd = &cobra.Command{
	Use:     "delete <tag>",
	Short:   "Delete a release",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_deleteCmd).Standalone()

	release_deleteCmd.Flags().Bool("cleanup-tag", false, "Delete the specified tag in addition to its release")
	release_deleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	releaseCmd.AddCommand(release_deleteCmd)

	carapace.Gen(release_deleteCmd).PositionalCompletion(
		action.ActionReleases(release_deleteCmd),
	)
}
