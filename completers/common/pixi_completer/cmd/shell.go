package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:     "shell",
	Short:   "Start a shell in a pixi environment, run `exit` to leave the shell",
	Aliases: []string{"s"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellCmd).Standalone()

	shellCmd.Flags().Bool("as-is", false, "Shorthand for the combination of --no-install and --frozen")
	shellCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	shellCmd.Flags().String("change-ps1", "", "Do not change the PS1 variable when starting a prompt")
	shellCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	shellCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	shellCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	shellCmd.Flags().StringP("environment", "e", "", "The environment to activate in the shell")
	shellCmd.Flags().Bool("force-activate", false, "Do not use the environment activation cache. (default: true except in experimental mode)")
	shellCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lock file, doesn't update lock file if it isn't up-to-date with the manifest file")
	shellCmd.Flags().Bool("locked", false, "Check if lock file is up-to-date before installing the environment, aborts when lock file isn't up-to-date with the manifest file")
	shellCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	shellCmd.Flags().Bool("no-completions", false, "Do not source the autocompletion scripts from the environment")
	shellCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	shellCmd.Flags().Bool("no-hard-links", false, "Disallow hard links during package installation")
	shellCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock file")
	shellCmd.Flags().Bool("no-ref-links", false, "Disallow ref links (copy-on-write) during package installation")
	shellCmd.Flags().Bool("no-symbolic-links", false, "Disallow symbolic links during package installation")
	shellCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	shellCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	shellCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	shellCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	shellCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots) or 'system' (system store)")
	shellCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	shellCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	rootCmd.AddCommand(shellCmd)

	carapace.Gen(shellCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"change-ps1":            carapace.ActionValues("true", "false").StyleF(style.ForKeyword),
		"environment":           pixi.ActionEnvironments(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
