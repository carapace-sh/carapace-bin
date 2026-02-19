package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a conda package from a Pixi package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("as-is", false, "Shorthand for --no-install and --frozen")
	buildCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	buildCmd.Flags().StringP("build-dir", "b", "", "The directory to use for incremental builds artifacts")
	buildCmd.Flags().String("build-platform", "", "The build platform to use for building")
	buildCmd.Flags().BoolP("clean", "c", false, "Whether to clean the build directory before building")
	buildCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	buildCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	buildCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile")
	buildCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment")
	buildCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	buildCmd.Flags().Bool("no-lockfile-update", false, "Skips lock-file updates")
	buildCmd.Flags().StringP("output-dir", "o", "", "The output directory to place the built artifacts")
	buildCmd.Flags().String("path", "", "The path to a directory containing a package manifest")
	buildCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	buildCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	buildCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	buildCmd.Flags().StringP("target-platform", "t", "", "The target platform to build for")
	buildCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	buildCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	buildCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"build-dir":             carapace.ActionDirectories(),
		"output-dir":            carapace.ActionDirectories(),
		"path":                  carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
