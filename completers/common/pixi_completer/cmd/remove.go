package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Removes dependencies from the workspace",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	removeCmd.Flags().String("branch", "", "The git branch")
	removeCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	removeCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	removeCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	removeCmd.Flags().StringP("feature", "f", "default", "The feature for which the dependency should be modified")
	removeCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lock file, doesn't update lock file if it isn't up-to-date with the manifest file")
	removeCmd.Flags().StringP("git", "g", "", "The git url to use when adding a git dependency")
	removeCmd.Flags().Bool("locked", false, "Check if lock file is up-to-date before installing the environment, aborts when lock file isn't up-to-date with the manifest file")
	removeCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	removeCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	removeCmd.Flags().Bool("no-hard-links", false, "Disallow hard links during package installation")
	removeCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock file")
	removeCmd.Flags().Bool("no-ref-links", false, "Disallow ref links (copy-on-write) during package installation")
	removeCmd.Flags().Bool("no-symbolic-links", false, "Disallow symbolic links during package installation")
	removeCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	removeCmd.Flags().StringSliceP("platform", "p", nil, "The platform for which the dependency should be modified. Must be the name of a platform already defined in the workspace")
	removeCmd.Flags().Bool("pypi", false, "The specified dependencies are pypi dependencies. Conflicts with `host` and `build`")
	removeCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	removeCmd.Flags().String("rev", "", "The git revision")
	removeCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	removeCmd.Flags().StringP("subdir", "s", "", "The subdirectory of the git repository to use")
	removeCmd.Flags().String("tag", "", "The git tag")
	removeCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	removeCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots) or 'system' (system store)")
	removeCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	removeCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
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
