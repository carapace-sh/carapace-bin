package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds dependencies to the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	addCmd.Flags().String("branch", "", "The git branch")
	addCmd.Flags().Bool("build", false, "The specified dependencies are build dependencies")
	addCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	addCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	addCmd.Flags().Bool("editable", false, "Whether the pypi requirement should be editable")
	addCmd.Flags().StringP("feature", "f", "", "The feature for which the dependency should be modified")
	addCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile")
	addCmd.Flags().StringP("git", "g", "", "The git url to use when adding a git dependency")
	addCmd.Flags().Bool("host", false, "The specified dependencies are host dependencies")
	addCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment")
	addCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	addCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	addCmd.Flags().Bool("no-lockfile-update", false, "Skips lock-file updates")
	addCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	addCmd.Flags().StringP("platform", "p", "", "The platform for which the dependency should be modified")
	addCmd.Flags().Bool("pypi", false, "The specified dependencies are pypi dependencies")
	addCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	addCmd.Flags().String("rev", "", "The git revision")
	addCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	addCmd.Flags().StringP("subdir", "s", "", "The subdirectory of the git repository to use")
	addCmd.Flags().String("tag", "", "The git tag")
	addCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	addCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	addCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"feature":               pixi.ActionFeatures(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"platform":              pixi.ActionPlatforms(),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})

	carapace.Gen(addCmd).PositionalAnyCompletion(
		pixi.ActionPackageSearch(),
	)
}
