package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_expose_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove exposed binaries from the global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_expose_removeCmd).Standalone()

	global_expose_removeCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	global_expose_removeCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	global_expose_removeCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	global_expose_removeCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	global_expose_removeCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	global_expose_removeCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	global_expose_removeCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	global_expose_removeCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	global_expose_removeCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	global_exposeCmd.AddCommand(global_expose_removeCmd)

	carapace.Gen(global_expose_removeCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
