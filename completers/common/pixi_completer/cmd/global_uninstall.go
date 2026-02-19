package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var global_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls environments from the global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_uninstallCmd).Standalone()

	global_uninstallCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	global_uninstallCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	global_uninstallCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	global_uninstallCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	global_uninstallCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	global_uninstallCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	global_uninstallCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	global_uninstallCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	global_uninstallCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	globalCmd.AddCommand(global_uninstallCmd)

	carapace.Gen(global_uninstallCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})

	carapace.Gen(global_uninstallCmd).PositionalAnyCompletion(
		pixi.ActionGlobalEnvironments(),
	)
}
