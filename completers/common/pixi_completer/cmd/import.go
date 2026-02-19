package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports a file into an environment in an existing workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	importCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	importCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	importCmd.Flags().StringP("environment", "e", "", "The environment to install")
	importCmd.Flags().StringP("feature", "f", "", "The feature for which the dependency should be modified")
	importCmd.Flags().String("format", "", "The format of the file to import")
	importCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	importCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	importCmd.Flags().StringP("platform", "p", "", "The platform for which the dependency should be modified")
	importCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	importCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	importCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	importCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	importCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
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
