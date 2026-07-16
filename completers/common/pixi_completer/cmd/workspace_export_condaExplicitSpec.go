package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_export_condaExplicitSpecCmd = &cobra.Command{
	Use:     "conda-explicit-spec",
	Short:   "Export workspace environment to a conda explicit specification file",
	Aliases: []string{"ces"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_export_condaExplicitSpecCmd).Standalone()

	workspace_export_condaExplicitSpecCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	workspace_export_condaExplicitSpecCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	workspace_export_condaExplicitSpecCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	workspace_export_condaExplicitSpecCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	workspace_export_condaExplicitSpecCmd.Flags().StringSliceP("environment", "e", nil, "The environments to render. Can be repeated for multiple environments")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lock file, doesn't update lock file if it isn't up-to-date with the manifest file")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("ignore-pypi-errors", false, "PyPI dependencies are not supported in the conda explicit spec file")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("ignore-source-errors", false, "Source dependencies are not supported in the conda explicit spec file")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("locked", false, "Check if lock file is up-to-date before installing the environment, aborts when lock file isn't up-to-date with the manifest file")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("no-hard-links", false, "Disallow hard links during package installation")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock file")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("no-ref-links", false, "Disallow ref links (copy-on-write) during package installation")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("no-symbolic-links", false, "Disallow symbolic links during package installation")
	workspace_export_condaExplicitSpecCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	workspace_export_condaExplicitSpecCmd.Flags().StringSliceP("platform", "p", nil, "The platform to render. Can be repeated for multiple platforms. Defaults to all platforms available for selected environments")
	workspace_export_condaExplicitSpecCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	workspace_export_condaExplicitSpecCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots) or 'system' (system store)")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	workspace_exportCmd.AddCommand(workspace_export_condaExplicitSpecCmd)

	carapace.Gen(workspace_export_condaExplicitSpecCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionEnvironments(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"platform":              pixi.ActionPlatforms(),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
