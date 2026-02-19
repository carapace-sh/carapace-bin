package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_export_condaExplicitSpecCmd = &cobra.Command{
	Use:   "conda-explicit-spec",
	Short: "Export workspace environment to a conda explicit specification file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_export_condaExplicitSpecCmd).Standalone()

	workspace_export_condaExplicitSpecCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	workspace_export_condaExplicitSpecCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	workspace_export_condaExplicitSpecCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	workspace_export_condaExplicitSpecCmd.Flags().StringP("environment", "e", "", "The environment to export")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("ignore-pypi-errors", false, "Ignore PyPI errors")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("ignore-source-errors", false, "Ignore source errors")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment")
	workspace_export_condaExplicitSpecCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("no-lockfile-update", false, "Skips lock-file updates")
	workspace_export_condaExplicitSpecCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	workspace_export_condaExplicitSpecCmd.Flags().StringP("platform", "p", "", "The platform to export")
	workspace_export_condaExplicitSpecCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	workspace_export_condaExplicitSpecCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	workspace_export_condaExplicitSpecCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	workspace_exportCmd.AddCommand(workspace_export_condaExplicitSpecCmd)

	carapace.Gen(workspace_export_condaExplicitSpecCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"environment":           pixi.ActionEnvironments(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"platform":              pixi.ActionPlatforms(),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
