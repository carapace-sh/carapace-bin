package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var plugin_createCmd = &cobra.Command{
	Use:   "create [OPTIONS] PLUGIN PLUGIN-DATA-DIR",
	Short: "Create a plugin from a rootfs and configuration. Plugin data directory must contain config.json and rootfs directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_createCmd).Standalone()

	plugin_createCmd.Flags().Bool("compress", false, "Compress the context using gzip")
	pluginCmd.AddCommand(plugin_createCmd)

	carapace.Gen(plugin_createCmd).PositionalCompletion(
		docker.ActionPlugins(),
		carapace.ActionDirectories(),
	)
}
