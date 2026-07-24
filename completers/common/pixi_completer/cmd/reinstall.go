package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "Re-install an environment, both updating the lock file and re-installing the environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	reinstallCmd.Flags().BoolP("all", "a", false, "Install all environments")
	reinstallCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	reinstallCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	reinstallCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	reinstallCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	reinstallCmd.Flags().StringSliceP("environment", "e", nil, "The environment to install")
	reinstallCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lock file, doesn't update lock file if it isn't up-to-date with the manifest file")
	reinstallCmd.Flags().Bool("locked", false, "Check if lock file is up-to-date before installing the environment, aborts when lock file isn't up-to-date with the manifest file")
	reinstallCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	reinstallCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	reinstallCmd.Flags().Bool("no-hard-links", false, "Disallow hard links during package installation")
	reinstallCmd.Flags().Bool("no-ref-links", false, "Disallow ref links (copy-on-write) during package installation")
	reinstallCmd.Flags().Bool("no-symbolic-links", false, "Disallow symbolic links during package installation")
	reinstallCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	reinstallCmd.Flags().StringP("platform", "p", "", "Re-install for the given platform; a warning is printed when it doesn't run on this machine")
	reinstallCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	reinstallCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	reinstallCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	reinstallCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots) or 'system' (system store)")
	reinstallCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	reinstallCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	rootCmd.AddCommand(reinstallCmd)

	carapace.Gen(reinstallCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionEnvironments(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
