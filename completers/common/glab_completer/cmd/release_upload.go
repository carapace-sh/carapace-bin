package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_uploadCmd = &cobra.Command{
	Use:   "upload <tag> [<files>...]",
	Short: "Upload release asset files or links to a GitLab release.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_uploadCmd).Standalone()

	release_uploadCmd.Flags().StringP("assets-links", "a", "", "`JSON` string representation of assets links, like: `--assets-links='[{\"name\": \"Asset1\", \"url\":\"https://<domain>/some/location/1\", \"link_type\": \"other\", \"direct_asset_path\": \"path/to/file\"}]'.`")
	releaseCmd.AddCommand(release_uploadCmd)

	carapace.Gen(release_uploadCmd).PositionalCompletion(
		action.ActionReleases(release_uploadCmd),
	)

	carapace.Gen(release_uploadCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
