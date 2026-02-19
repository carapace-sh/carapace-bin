package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs task in the pixi environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("as-is", false, "Shorthand for --no-install and --frozen")
	runCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	runCmd.Flags().Bool("clean-env", false, "Use a clean environment")
	runCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	runCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	runCmd.Flags().BoolP("dry-run", "n", false, "Don't actually run the task")
	runCmd.Flags().StringP("environment", "e", "", "The environment to run in")
	runCmd.Flags().Bool("force-activate", false, "Force activation of the environment")
	runCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile")
	runCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment")
	runCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	runCmd.Flags().Bool("no-completions", false, "Disable completions")
	runCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	runCmd.Flags().Bool("no-lockfile-update", false, "Skips lock-file updates")
	runCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	runCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	runCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	runCmd.Flags().Bool("skip-deps", false, "Skip running the task dependencies")
	runCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	runCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	runCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionEnvironments(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		pixi.ActionTasks(),
	)
}
