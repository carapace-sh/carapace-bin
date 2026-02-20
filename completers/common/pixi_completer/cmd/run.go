package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Runs task in the pixi environment",
	Aliases: []string{"r"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("as-is", false, "Shorthand for the combination of --no-install and --frozen")
	runCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	runCmd.Flags().Bool("clean-env", false, "Use a clean environment to run the task")
	runCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	runCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	runCmd.Flags().BoolP("dry-run", "n", false, "Run the task in dry-run mode (only print the command that would run)")
	runCmd.Flags().StringP("environment", "e", "", "The environment to run the task in")
	runCmd.Flags().Bool("force-activate", false, "Do not use the environment activation cache. (default: true except in experimental mode)")
	runCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile, doesn't update lockfile if it isn't up-to-date with the manifest file")
	runCmd.Flags().BoolS("h", "h", false, "")
	runCmd.Flags().Bool("help", false, "")
	runCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment, aborts when lockfile isn't up-to-date with the manifest file")
	runCmd.Flags().Bool("no-completions", false, "Do not source the autocompletion scripts from the environment")
	runCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	runCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	runCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	runCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	runCmd.Flags().Bool("skip-deps", false, "Don't run the dependencies of the task ('depends-on' field in the task definition)")
	runCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	runCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	runCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionEnvironments(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		pixi.ActionTasks(),
	)
}
