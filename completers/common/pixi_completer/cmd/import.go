package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports a file into an environment in an existing workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	importCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	importCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	importCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	importCmd.Flags().StringP("environment", "e", "", "A name for the created environment")
	importCmd.Flags().StringP("feature", "f", "", "A name for the created feature")
	importCmd.Flags().String("format", "", "Which format to interpret the file as")
	importCmd.PersistentFlags().StringP("manifest-path", "m", "", "The path to `pixi.toml`, `pyproject.toml`, or the workspace directory")
	importCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	importCmd.Flags().Bool("no-hard-links", false, "Disallow hard links during package installation")
	importCmd.Flags().Bool("no-ref-links", false, "Disallow ref links (copy-on-write) during package installation")
	importCmd.Flags().Bool("no-symbolic-links", false, "Disallow symbolic links during package installation")
	importCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	importCmd.Flags().StringSliceP("platform", "p", nil, "The platforms for the imported environment. Accepts a workspace platform name; a bare conda subdir (e.g. `linux-64`) is also accepted. Names that aren't yet declared get auto-added as subdir platforms")
	importCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	importCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	importCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	importCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots) or 'system' (system store)")
	importCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	importCmd.PersistentFlags().StringP("workspace", "w", "", "Name of the workspace")
	rootCmd.AddCommand(importCmd)

	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionEnvironments(),
		"feature":               pixi.ActionFeatures(),
		"format":                carapace.ActionValues("conda-env", "pypi-txt"),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"platform":              pixi.ActionPlatforms(),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
