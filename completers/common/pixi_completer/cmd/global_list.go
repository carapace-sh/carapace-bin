package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var global_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Lists global environments with their dependencies and exposed commands. Can also display all packages within a specific global environment when using the --environment flag.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_listCmd).Standalone()

	global_listCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	global_listCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	global_listCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	global_listCmd.Flags().StringP("environment", "e", "", "Allows listing all the packages installed in a specific environment, with an output similar to `pixi list`")
	global_listCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	global_listCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	global_listCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	global_listCmd.Flags().String("sort-by", "", "Sorting strategy for the package table of an environment")
	global_listCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	global_listCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	global_listCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	globalCmd.AddCommand(global_listCmd)

	carapace.Gen(global_listCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionGlobalEnvironments(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
		"sort-by":               carapace.ActionValues("size", "name"),
	})
}
