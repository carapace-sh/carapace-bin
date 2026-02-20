package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var global_addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Adds dependencies to an environment",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_addCmd).Standalone()

	global_addCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	global_addCmd.Flags().String("branch", "", "The git branch")
	global_addCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	global_addCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	global_addCmd.Flags().StringP("environment", "e", "", "Specifies the environment that the dependencies need to be added to")
	global_addCmd.Flags().StringSlice("expose", nil, "Add one or more mapping which describe which executables are exposed. The syntax is `exposed_name=executable_name`, so for example `python3.10=python`. Alternatively, you can input only an executable_name and `executable_name=executable_name` is assumed")
	global_addCmd.Flags().String("git", "", "The git url, e.g. `https://github.com/user/repo.git`")
	global_addCmd.Flags().String("path", "", "The path to the local package")
	global_addCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	global_addCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	global_addCmd.Flags().String("rev", "", "The git revision")
	global_addCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	global_addCmd.Flags().String("subdir", "", "The subdirectory within the git repository")
	global_addCmd.Flags().String("tag", "", "The git tag")
	global_addCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	global_addCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	global_addCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	global_addCmd.MarkFlagRequired("environment")
	globalCmd.AddCommand(global_addCmd)

	carapace.Gen(global_addCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionGlobalEnvironments(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
