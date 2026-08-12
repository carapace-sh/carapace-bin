package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search a conda package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	searchCmd.Flags().StringSliceP("channel", "c", nil, "The channels to consider as a name or a url. Multiple channels can be specified by using this field multiple times")
	searchCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	searchCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	searchCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	searchCmd.Flags().Bool("json", false, "Output in JSON format")
	searchCmd.Flags().StringP("limit", "l", "5", "Limit the number of versions shown per package, -1 for no limit")
	searchCmd.Flags().StringP("limit-packages", "n", "5", "Limit the number of packages shown, -1 for no limit")
	searchCmd.Flags().Bool("offline", false, "Run without network access, using only cached data. Commands fail if data is missing from the cache. Pass `--offline=false` to override an `offline` setting from the configuration")
	searchCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	searchCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	searchCmd.Flags().Bool("no-hard-links", false, "Disallow hard links during package installation")
	searchCmd.Flags().Bool("no-ref-links", false, "Disallow ref links (copy-on-write) during package installation")
	searchCmd.Flags().Bool("no-symbolic-links", false, "Disallow symbolic links during package installation")
	searchCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	searchCmd.Flags().StringP("platform", "p", "", "The platform to search packages for. By default, searches all platforms from the manifest (or all known platforms if no manifest is found). Accepts a workspace platform name; a bare conda subdir (e.g. `linux-64`) is also accepted")
	searchCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	searchCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	searchCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	searchCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots) or 'system' (system store)")
	searchCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	searchCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"config-file":           carapace.ActionFiles(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"platform":              pixi.ActionPlatforms(),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
		"tls-root-certs":        carapace.ActionValues("webpki", "system"),
	})
}
