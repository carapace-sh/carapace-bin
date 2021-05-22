package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_downloadCmd = &cobra.Command{
	Use:   "download [<tag>]",
	Short: "Download release assets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	release_downloadCmd.Flags().StringP("dir", "D", ".", "The directory to download files into")
	release_downloadCmd.Flags().StringArrayP("pattern", "p", nil, "Download only assets that match a glob pattern")
	releaseCmd.AddCommand(release_downloadCmd)

	carapace.Gen(release_downloadCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
		"pattern": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionReleaseAssets(release_downloadCmd, c.Args[0])
			} else {
				return carapace.ActionValues()
			}
		}),
	})

	carapace.Gen(release_downloadCmd).PositionalCompletion(
		action.ActionReleases(release_downloadCmd),
	)

}
