package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start a shell in a pixi environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellCmd).Standalone()

	shellCmd.Flags().Bool("as-is", false, "Shorthand for --no-install and --frozen")
	shellCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	shellCmd.Flags().String("change-ps1", "", "Whether to change the PS1")
	shellCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	shellCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	shellCmd.Flags().StringP("environment", "e", "", "The environment to activate")
	shellCmd.Flags().Bool("force-activate", false, "Force activation of the environment")
	shellCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile")
	shellCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment")
	shellCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	shellCmd.Flags().Bool("no-completions", false, "Disable completions")
	shellCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	shellCmd.Flags().Bool("no-lockfile-update", false, "Skips lock-file updates")
	shellCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	shellCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	shellCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	shellCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	shellCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	shellCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	rootCmd.AddCommand(shellCmd)

	carapace.Gen(shellCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"change-ps1":            carapace.ActionValues("true", "false"),
		"environment":           pixi.ActionEnvironments(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
