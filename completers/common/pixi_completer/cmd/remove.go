package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes dependencies from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	removeCmd.Flags().String("branch", "", "The git branch")
	removeCmd.Flags().Bool("build", false, "The specified dependencies are build dependencies")
	removeCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	removeCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	removeCmd.Flags().StringP("feature", "f", "", "The feature for which the dependency should be modified")
	removeCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile")
	removeCmd.Flags().StringP("git", "g", "", "The git url to use when adding a git dependency")
	removeCmd.Flags().Bool("host", false, "The specified dependencies are host dependencies")
	removeCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment")
	removeCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	removeCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	removeCmd.Flags().Bool("no-lockfile-update", false, "Skips lock-file updates")
	removeCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	removeCmd.Flags().StringP("platform", "p", "", "The platform for which the dependency should be modified")
	removeCmd.Flags().Bool("pypi", false, "The specified dependencies are pypi dependencies")
	removeCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	removeCmd.Flags().String("rev", "", "The git revision")
	removeCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	removeCmd.Flags().StringP("subdir", "s", "", "The subdirectory of the git repository to use")
	removeCmd.Flags().String("tag", "", "The git tag")
	removeCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	removeCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	removeCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"feature":               pixi.ActionFeatures(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"platform":              pixi.ActionPlatforms(),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		pixi.ActionPackages(),
	)
}
