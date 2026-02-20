package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "The `update` command checks if there are newer versions of the dependencies and updates the `pixi.lock` file and environments accordingly",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	updateCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	updateCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	updateCmd.Flags().BoolP("dry-run", "n", false, "Don't actually write the lockfile or update any environment")
	updateCmd.Flags().StringSliceP("environment", "e", nil, "The environments to update. If none is specified, all environments are updated")
	updateCmd.Flags().Bool("json", false, "Output the changes in JSON format")
	updateCmd.Flags().Bool("no-install", false, "Don't install the (solve) environments needed for pypi-dependencies solving")
	updateCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	updateCmd.Flags().StringSliceP("platform", "p", nil, "The platforms to update. If none is specified, all platforms are updated")
	updateCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	updateCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	updateCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	updateCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	updateCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionEnvironments(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"platform":              pixi.ActionPlatforms(),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
