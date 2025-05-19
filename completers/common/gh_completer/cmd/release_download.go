package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_downloadCmd = &cobra.Command{
	Use:     "download [<tag>]",
	Short:   "Download release assets",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_downloadCmd).Standalone()

	release_downloadCmd.Flags().StringP("archive", "A", "", "Download the source code archive in the specified `format` (zip or tar.gz)")
	release_downloadCmd.Flags().Bool("clobber", false, "Overwrite existing files of the same name")
	release_downloadCmd.Flags().StringP("dir", "D", "", "The `directory` to download files into")
	release_downloadCmd.Flags().StringP("output", "O", "", "The `file` to write a single asset to (use \"-\" to write to standard output)")
	release_downloadCmd.Flags().StringSliceP("pattern", "p", nil, "Download only assets that match a glob pattern")
	release_downloadCmd.Flags().Bool("skip-existing", false, "Skip downloading when files of the same name exist")
	releaseCmd.AddCommand(release_downloadCmd)

	carapace.Gen(release_downloadCmd).FlagCompletion(carapace.ActionMap{
		"archive": carapace.ActionValues("zip", "tar.gz"),
		"dir":     carapace.ActionDirectories(),
		"output": carapace.Batch(
			carapace.ActionValuesDescribed("-", "standard output"),
			carapace.ActionFiles(),
		).ToA(),
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
