package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync global manifest with installed environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_syncCmd).Standalone()

	global_syncCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	global_syncCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	global_syncCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	global_syncCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	global_syncCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	global_syncCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	global_syncCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	global_syncCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	global_syncCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	globalCmd.AddCommand(global_syncCmd)

	carapace.Gen(global_syncCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
