package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_downloadCmd = &cobra.Command{
	Use:   "download <tag>",
	Short: "Download asset files from a GitLab release.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_downloadCmd).Standalone()

	release_downloadCmd.Flags().StringSliceP("asset-name", "n", nil, "Download only assets that match the name or a glob pattern.")
	release_downloadCmd.Flags().StringP("dir", "D", "", "Directory to download the release assets to.")
	releaseCmd.AddCommand(release_downloadCmd)

	carapace.Gen(release_downloadCmd).FlagCompletion(carapace.ActionMap{
		"asset-name": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionReleaseAssets(release_downloadCmd, c.Args[0]).UniqueList(",")
			}
			return carapace.ActionValues()
		}),
		"dir": carapace.ActionDirectories(),
	})

	carapace.Gen(release_downloadCmd).PositionalCompletion(
		action.ActionReleases(release_downloadCmd),
	)
}
