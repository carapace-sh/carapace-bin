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
	release_uploadCmd.Flags().String("package-name", "", "The package name to use when uploading the assets to the generic package release with --use-package-registry.")
	release_uploadCmd.Flags().Bool("use-package-registry", false, "Upload release assets to the generic package registry of the project. Alternatively to this flag you may also set the GITLAB_RELEASE_ASSETS_USE_PACKAGE_REGISTRY environment variable to either the value true or 1. The flag takes precedence over this environment variable.")
	releaseCmd.AddCommand(release_uploadCmd)

	carapace.Gen(release_uploadCmd).PositionalCompletion(
		action.ActionReleases(release_uploadCmd),
	)

	carapace.Gen(release_uploadCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
