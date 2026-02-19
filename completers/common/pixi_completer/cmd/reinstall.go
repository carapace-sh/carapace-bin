package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "Re-install an environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	reinstallCmd.Flags().BoolP("all", "a", false, "Reinstall all environments")
	reinstallCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	reinstallCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	reinstallCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	reinstallCmd.Flags().StringP("environment", "e", "", "The environment to install")
	reinstallCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile")
	reinstallCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment")
	reinstallCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	reinstallCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	reinstallCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	reinstallCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	reinstallCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	reinstallCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	reinstallCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	rootCmd.AddCommand(reinstallCmd)

	carapace.Gen(reinstallCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionEnvironments(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
