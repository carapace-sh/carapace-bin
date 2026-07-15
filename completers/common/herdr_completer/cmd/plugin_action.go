package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_actionCmd = &cobra.Command{
	Use:   "action",
	Short: "List or invoke plugin actions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_actionCmd).Standalone()

	pluginCmd.AddCommand(plugin_actionCmd)
}
