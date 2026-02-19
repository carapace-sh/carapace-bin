package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var global_expose_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add exposed binaries from an environment to your global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_expose_addCmd).Standalone()

	global_expose_addCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	global_expose_addCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	global_expose_addCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	global_expose_addCmd.Flags().StringP("environment", "e", "", "Specifies the environment that the dependencies need to be added to")
	global_expose_addCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	global_expose_addCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	global_expose_addCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	global_expose_addCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	global_expose_addCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	global_expose_addCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	global_exposeCmd.AddCommand(global_expose_addCmd)

	carapace.Gen(global_expose_addCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionGlobalEnvironments(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
