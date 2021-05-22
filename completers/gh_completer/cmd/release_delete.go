package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_deleteCmd = &cobra.Command{
	Use:   "delete <tag>",
	Short: "Delete a release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	release_deleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	releaseCmd.AddCommand(release_deleteCmd)

	carapace.Gen(release_deleteCmd).PositionalCompletion(
		action.ActionReleases(release_deleteCmd),
	)
}
