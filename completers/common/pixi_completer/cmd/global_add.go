package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var global_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds dependencies to an environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_addCmd).Standalone()

	global_addCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	global_addCmd.Flags().String("branch", "", "The git branch")
	global_addCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	global_addCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	global_addCmd.Flags().StringP("environment", "e", "", "Specifies the environment that the dependencies need to be added to")
	global_addCmd.Flags().String("expose", "", "Add one or more mapping which describe which executables are exposed")
	global_addCmd.Flags().String("git", "", "The git url")
	global_addCmd.Flags().String("path", "", "The path to the local package")
	global_addCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	global_addCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	global_addCmd.Flags().String("rev", "", "The git revision")
	global_addCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	global_addCmd.Flags().String("subdir", "", "The subdirectory within the git repository")
	global_addCmd.Flags().String("tag", "", "The git tag")
	global_addCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	global_addCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	global_addCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	globalCmd.AddCommand(global_addCmd)

	carapace.Gen(global_addCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionGlobalEnvironments(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
