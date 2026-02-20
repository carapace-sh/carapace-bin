package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install an environment, both updating the lockfile and installing the environment",
	Aliases: []string{"i"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().BoolP("all", "a", false, "Install all environments")
	installCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	installCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	installCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	installCmd.Flags().StringSliceP("environment", "e", nil, "The environment to install")
	installCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile, doesn't update lockfile if it isn't up-to-date with the manifest file")
	installCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment, aborts when lockfile isn't up-to-date with the manifest file")
	installCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	installCmd.Flags().StringSlice("only", nil, "Install and build only these package(s) and their dependencies. Can be passed multiple times")
	installCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	installCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	installCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	installCmd.Flags().StringSlice("skip", nil, "Skip installation of specific packages present in the lockfile. This uses a soft exclusion: the package will be skipped but its dependencies are installed")
	installCmd.Flags().StringSlice("skip-with-deps", nil, "Skip a package and its entire dependency subtree. This performs a hard exclusion: the package and its dependencies are not installed unless reachable from another non-skipped root")
	installCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	installCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	installCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionEnvironments(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
