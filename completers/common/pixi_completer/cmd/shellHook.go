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

	shellHookCmd.Flags().Bool("as-is", false, "Shorthand for the combination of --no-install and --frozen")
	shellHookCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	shellHookCmd.Flags().String("change-ps1", "", "Do not change the PS1 variable when starting a prompt")
	shellHookCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	shellHookCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	shellHookCmd.Flags().StringP("environment", "e", "", "The environment to activate in the script")
	shellHookCmd.Flags().Bool("force-activate", false, "Do not use the environment activation cache. (default: true except in experimental mode)")
	shellHookCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile, doesn't update lockfile if it isn't up-to-date with the manifest file")
	shellHookCmd.Flags().Bool("json", false, "Emit the environment variables set by running the activation as JSON")
	shellHookCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment, aborts when lockfile isn't up-to-date with the manifest file")
	shellHookCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	shellHookCmd.Flags().Bool("no-completions", false, "Do not source the autocompletion scripts from the environment")
	shellHookCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	shellHookCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	shellHookCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	shellHookCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	shellHookCmd.Flags().StringP("shell", "s", "", "Sets the shell, options: [`bash`,  `zsh`,  `xonsh`,  `cmd`, `powershell`,  `fish`,  `nushell`]")
	shellHookCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	shellHookCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	shellHookCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
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
