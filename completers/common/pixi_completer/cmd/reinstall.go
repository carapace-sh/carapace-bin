package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "Re-install an environment, both updating the lockfile and re-installing the environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	reinstallCmd.Flags().BoolP("all", "a", false, "Install all environments")
	reinstallCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	reinstallCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	reinstallCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	reinstallCmd.Flags().StringSliceP("environment", "e", nil, "The environment to install")
	reinstallCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile, doesn't update lockfile if it isn't up-to-date with the manifest file")
	reinstallCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment, aborts when lockfile isn't up-to-date with the manifest file")
	reinstallCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	reinstallCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	reinstallCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	reinstallCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	reinstallCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	reinstallCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	rootCmd.AddCommand(reinstallCmd)

	carapace.Gen(reinstallCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionEnvironments(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
