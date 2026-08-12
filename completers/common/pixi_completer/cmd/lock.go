package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Solve environment and update the lock file without installing the environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lockCmd).Standalone()
	lockCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	lockCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	lockCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	lockCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")

	lockCmd.Flags().Bool("check", false, "Check if any changes have been made to the lock file. If yes, exit with a non-zero code")
	lockCmd.Flags().Bool("dry-run", false, "Compute the lock file without writing to disk. Implies --no-install")
	lockCmd.Flags().Bool("json", false, "Output the changes in JSON format")
	lockCmd.Flags().Bool("offline", false, "Run without network access, using only cached data. Commands fail if data is missing from the cache. Pass `--offline=false` to override an `offline` setting from the configuration")
	lockCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	lockCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	lockCmd.Flags().Bool("no-hard-links", false, "Disallow hard links during package installation")
	lockCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock file")
	lockCmd.Flags().Bool("no-ref-links", false, "Disallow ref links (copy-on-write) during package installation")
	lockCmd.Flags().Bool("no-symbolic-links", false, "Disallow symbolic links during package installation")
	lockCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	lockCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	lockCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	lockCmd.PersistentFlags().StringP("script", "s", "", "The path to a Python script containing PEP 723 metadata. Pixi run also accepts an HTTP(S) URL or '-' to read the script from stdin")
	lockCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	lockCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots) or 'system' (system store)")
	lockCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	lockCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	rootCmd.AddCommand(lockCmd)

	carapace.Gen(lockCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"config-file":           carapace.ActionFiles(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
		"script":                carapace.ActionFiles(".py"),
		"tls-root-certs":        carapace.ActionValues("webpki", "system"),
	})
}
