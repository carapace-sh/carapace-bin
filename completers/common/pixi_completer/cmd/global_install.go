package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var global_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs the defined packages in a globally accessible location",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_installCmd).Standalone()

	global_installCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	global_installCmd.Flags().String("branch", "", "The git branch")
	global_installCmd.Flags().StringP("channel", "c", "", "The channels to consider as a name or a url")
	global_installCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	global_installCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	global_installCmd.Flags().StringP("environment", "e", "", "Ensures that all packages will be installed in the same environment")
	global_installCmd.Flags().String("expose", "", "Add one or more mapping which describe which executables are exposed")
	global_installCmd.Flags().Bool("force-reinstall", false, "Specifies that the environment should be reinstalled")
	global_installCmd.Flags().String("git", "", "The git url")
	global_installCmd.Flags().Bool("no-shortcuts", false, "Specifies that no shortcuts should be created")
	global_installCmd.Flags().String("path", "", "The path to the local package")
	global_installCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	global_installCmd.Flags().StringP("platform", "p", "", "The platform to install the packages for")
	global_installCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	global_installCmd.Flags().String("rev", "", "The git revision")
	global_installCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	global_installCmd.Flags().String("subdir", "", "The subdirectory within the git repository")
	global_installCmd.Flags().String("tag", "", "The git tag")
	global_installCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	global_installCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	global_installCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	global_installCmd.Flags().String("with", "", "Add additional dependencies to the environment")
	globalCmd.AddCommand(global_installCmd)

	carapace.Gen(global_installCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionGlobalEnvironments(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
