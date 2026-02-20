package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports a file into an environment in an existing workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	importCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	importCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	importCmd.Flags().StringP("environment", "e", "", "A name for the created environment")
	importCmd.Flags().StringP("feature", "f", "", "A name for the created feature")
	importCmd.Flags().String("format", "", "Which format to interpret the file as")
	importCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	importCmd.Flags().StringSliceP("platform", "p", nil, "The platforms for the imported environment")
	importCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	importCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	importCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	importCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	importCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	rootCmd.AddCommand(importCmd)

	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionEnvironments(),
		"feature":               pixi.ActionFeatures(),
		"format":                carapace.ActionValues("conda-env", "pypi-txt"),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"platform":              pixi.ActionPlatforms(),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
