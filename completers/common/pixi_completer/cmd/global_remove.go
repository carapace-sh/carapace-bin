package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var global_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Removes dependencies from an environment",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_removeCmd).Standalone()

	global_removeCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	global_removeCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	global_removeCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	global_removeCmd.Flags().StringP("environment", "e", "", "Specifies the environment that the dependencies need to be removed from")
	global_removeCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	global_removeCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	global_removeCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	global_removeCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	global_removeCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	global_removeCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	globalCmd.AddCommand(global_removeCmd)

	carapace.Gen(global_removeCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionGlobalEnvironments(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
