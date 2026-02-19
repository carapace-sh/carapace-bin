package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "The update command checks if there are newer versions of the dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	updateCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	updateCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	updateCmd.Flags().BoolP("dry-run", "n", false, "Don't actually update")
	updateCmd.Flags().StringP("environment", "e", "", "The environment to update")
	updateCmd.Flags().Bool("json", false, "Whether to show the output as JSON or not")
	updateCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	updateCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	updateCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	updateCmd.Flags().StringP("platform", "p", "", "The platform to update")
	updateCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	updateCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	updateCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	updateCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	updateCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionEnvironments(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"platform":              pixi.ActionPlatforms(),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
