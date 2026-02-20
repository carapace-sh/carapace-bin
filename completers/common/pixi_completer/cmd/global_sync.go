package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_syncCmd = &cobra.Command{
	Use:     "sync",
	Short:   "Sync global manifest with installed environments",
	Aliases: []string{"s"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_syncCmd).Standalone()

	global_syncCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	global_syncCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	global_syncCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	global_syncCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	global_syncCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	global_syncCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	global_syncCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	global_syncCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	global_syncCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	globalCmd.AddCommand(global_syncCmd)

	carapace.Gen(global_syncCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
