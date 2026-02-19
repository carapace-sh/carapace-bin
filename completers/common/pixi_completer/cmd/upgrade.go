package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Checks if there are newer versions of the dependencies and upgrades them",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	upgradeCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	upgradeCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	upgradeCmd.Flags().BoolP("dry-run", "n", false, "Don't actually upgrade")
	upgradeCmd.Flags().String("exclude", "", "Packages to exclude from the upgrade")
	upgradeCmd.Flags().StringP("feature", "f", "", "The feature for which the dependency should be modified")
	upgradeCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile")
	upgradeCmd.Flags().Bool("json", false, "Whether to show the output as JSON or not")
	upgradeCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment")
	upgradeCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	upgradeCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	upgradeCmd.Flags().Bool("no-lockfile-update", false, "Skips lock-file updates")
	upgradeCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	upgradeCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	upgradeCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	upgradeCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	upgradeCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	upgradeCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"feature":               pixi.ActionFeatures(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		pixi.ActionPackages(),
	)
}
