package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var global_shortcut_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a shortcut from an environment to your machine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_shortcut_addCmd).Standalone()

	global_shortcut_addCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	global_shortcut_addCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	global_shortcut_addCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	global_shortcut_addCmd.Flags().StringP("environment", "e", "", "The environment from which the shortcut should be added")
	global_shortcut_addCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	global_shortcut_addCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	global_shortcut_addCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	global_shortcut_addCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	global_shortcut_addCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	global_shortcut_addCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	global_shortcut_addCmd.MarkFlagRequired("environment")
	global_shortcutCmd.AddCommand(global_shortcut_addCmd)

	carapace.Gen(global_shortcut_addCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionGlobalEnvironments(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
