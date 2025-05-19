package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_deleteAssetCmd = &cobra.Command{
	Use:     "delete-asset <tag> <asset-name>",
	Short:   "Delete an asset from a release",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_deleteAssetCmd).Standalone()

	release_deleteAssetCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	releaseCmd.AddCommand(release_deleteAssetCmd)

	carapace.Gen(release_deleteAssetCmd).PositionalCompletion(
		action.ActionReleases(release_deleteAssetCmd),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionReleaseAssets(release_deleteAssetCmd, c.Args[0])
		}),
	)
}
