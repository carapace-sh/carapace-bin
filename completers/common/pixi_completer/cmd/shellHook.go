package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var shellHookCmd = &cobra.Command{
	Use:   "shell-hook",
	Short: "Print the pixi environment activation script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellHookCmd).Standalone()

	shellHookCmd.Flags().Bool("as-is", false, "Shorthand for --no-install and --frozen")
	shellHookCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	shellHookCmd.Flags().String("change-ps1", "", "Whether to change the PS1")
	shellHookCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	shellHookCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	shellHookCmd.Flags().StringP("environment", "e", "", "The environment to activate")
	shellHookCmd.Flags().Bool("force-activate", false, "Force activation of the environment")
	shellHookCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile")
	shellHookCmd.Flags().Bool("json", false, "Whether to show the output as JSON or not")
	shellHookCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment")
	shellHookCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	shellHookCmd.Flags().Bool("no-completions", false, "Disable completions")
	shellHookCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	shellHookCmd.Flags().Bool("no-lockfile-update", false, "Skips lock-file updates")
	shellHookCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	shellHookCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	shellHookCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	shellHookCmd.Flags().StringP("shell", "s", "", "The shell to generate a hook for")
	shellHookCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	shellHookCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	shellHookCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	rootCmd.AddCommand(shellHookCmd)

	carapace.Gen(shellHookCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"change-ps1":            carapace.ActionValues("true", "false"),
		"environment":           pixi.ActionEnvironments(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
		"shell":                 carapace.ActionValues("bash", "zsh", "xonsh", "cmd", "powershell", "fish", "nushell"),
	})
}
