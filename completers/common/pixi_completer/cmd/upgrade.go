package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Checks if there are newer versions of the dependencies and upgrades them in the lockfile and manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	upgradeCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	upgradeCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	upgradeCmd.Flags().BoolP("dry-run", "n", false, "Only show the changes that would be made, without actually updating the manifest, lock file, or environment")
	upgradeCmd.Flags().StringSlice("exclude", nil, "The packages which should be excluded")
	upgradeCmd.Flags().StringP("feature", "f", "", "The feature to update")
	upgradeCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile, doesn't update lockfile if it isn't up-to-date with the manifest file")
	upgradeCmd.Flags().Bool("json", false, "Output the changes in JSON format")
	upgradeCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment, aborts when lockfile isn't up-to-date with the manifest file")
	upgradeCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	upgradeCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	upgradeCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	upgradeCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	upgradeCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	upgradeCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	upgradeCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	upgradeCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
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
