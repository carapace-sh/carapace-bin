package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_channel_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove channel(s) from the manifest and updates the lockfile",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_channel_removeCmd).Standalone()

	workspace_channel_removeCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	workspace_channel_removeCmd.Flags().String("concurrent-downloads", "", "Max concurrent network requests, default is `50`")
	workspace_channel_removeCmd.Flags().String("concurrent-solves", "", "Max concurrent solves, default is the number of CPUs")
	workspace_channel_removeCmd.Flags().StringP("feature", "f", "", "The name of the feature to modify")
	workspace_channel_removeCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile, doesn't update lockfile if it isn't up-to-date with the manifest file")
	workspace_channel_removeCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment, aborts when lockfile isn't up-to-date with the manifest file")
	workspace_channel_removeCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	workspace_channel_removeCmd.Flags().String("pinning-strategy", "", "Set pinning strategy")
	workspace_channel_removeCmd.Flags().Bool("prepend", false, "Add the channel(s) to the beginning of the channels list, making them the highest priority")
	workspace_channel_removeCmd.Flags().String("priority", "", "Specify the channel priority")
	workspace_channel_removeCmd.Flags().String("pypi-keyring-provider", "", "Specifies whether to use the keyring to look up credentials for PyPI")
	workspace_channel_removeCmd.Flags().Bool("run-post-link-scripts", false, "Run post-link scripts (insecure)")
	workspace_channel_removeCmd.Flags().Bool("tls-no-verify", false, "Do not verify the TLS certificate of the server")
	workspace_channel_removeCmd.Flags().String("tls-root-certs", "", "Which TLS root certificates to use: 'webpki' (bundled Mozilla roots), 'native' (system store), or 'all' (both)")
	workspace_channel_removeCmd.Flags().Bool("use-environment-activation-cache", false, "Use environment activation cache (experimental)")
	workspace_channelCmd.AddCommand(workspace_channel_removeCmd)

	carapace.Gen(workspace_channel_removeCmd).FlagCompletion(carapace.ActionMap{
		"auth-file":             carapace.ActionFiles(),
		"feature":               pixi.ActionFeatures(),
		"pinning-strategy":      carapace.ActionValues("semver", "minor", "major", "latest-up", "exact-version", "no-pin"),
		"pypi-keyring-provider": carapace.ActionValues("disabled", "subprocess"),
	})
}
