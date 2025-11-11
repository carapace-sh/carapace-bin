package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var plugin_upgradeCmd = &cobra.Command{
	Use:   "upgrade [OPTIONS] PLUGIN [REMOTE]",
	Short: "Upgrade an existing plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_upgradeCmd).Standalone()

	plugin_upgradeCmd.Flags().Bool("disable-content-trust", false, "Skip image verification (deprecated)")
	plugin_upgradeCmd.Flags().Bool("grant-all-permissions", false, "Grant all permissions necessary to run the plugin")
	plugin_upgradeCmd.Flags().Bool("skip-remote-check", false, "Do not check if specified remote plugin matches existing plugin image")
	plugin_upgradeCmd.Flag("disable-content-trust").Hidden = true
	pluginCmd.AddCommand(plugin_upgradeCmd)

	carapace.Gen(plugin_upgradeCmd).PositionalCompletion(
		docker.ActionPlugins(),
		carapace.ActionValues(),
	)
}
