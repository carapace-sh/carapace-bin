package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_channel_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove channel(s) from the manifest and updates the lockfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_channel_removeCmd).Standalone()

	workspace_channel_removeCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	workspace_channel_removeCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests")
	workspace_channel_removeCmd.Flags().String("concurrent-solves", "", "Max concurrent solves")
	workspace_channel_removeCmd.Flags().StringP("feature", "f", "", "The feature for which the dependency should be modified")
	workspace_channel_removeCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile")
	workspace_channel_removeCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment")
	workspace_channel_removeCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_channel_removeCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	workspace_channel_removeCmd.Flags().Bool("no-lockfile-update", false, "Skips lock-file updates")
	workspace_channel_removeCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	workspace_channel_removeCmd.Flags().Bool("prepend", false, "Prepend the channel to the list")
	workspace_channel_removeCmd.Flags().String("priority", "", "The priority of the channel to add")
	workspace_channel_removeCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring for PyPI")
	workspace_channel_removeCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts")
	workspace_channel_removeCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	workspace_channel_removeCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use")
	workspace_channel_removeCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache")
	workspace_channelCmd.AddCommand(workspace_channel_removeCmd)

	carapace.Gen(workspace_channel_removeCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"feature":               pixi.ActionFeatures(),
		"manifest-path":         carapace.ActionFiles(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
