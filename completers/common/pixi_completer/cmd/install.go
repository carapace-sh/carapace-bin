package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install an environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().BoolP("all", "a", false, "Install all environments")
	installCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	installCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	installCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	installCmd.Flags().StringP("environment", "e", "", "The environment to install")
	installCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile")
	installCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment")
	installCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	installCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	installCmd.Flags().String("only", "", "Only install certain environments")
	installCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	installCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	installCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	installCmd.Flags().String("skip", "", "Skip the installation of certain environments")
	installCmd.Flags().String("skip-with-deps", "", "Skip the installation of certain environments and their dependencies")
	installCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	installCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	installCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionEnvironments(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
