package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates environments in the global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_updateCmd).Standalone()

	global_updateCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	global_updateCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	global_updateCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	global_updateCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	global_updateCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	global_updateCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	global_updateCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	global_updateCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	global_updateCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	globalCmd.AddCommand(global_updateCmd)

	carapace.Gen(global_updateCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
