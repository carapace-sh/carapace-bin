package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove plugin and package versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_removeCmd).Standalone()

	pluginCmd.AddCommand(plugin_removeCmd)
}
