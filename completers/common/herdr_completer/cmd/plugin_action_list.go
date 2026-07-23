package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_action_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List plugin actions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_action_listCmd).Standalone()

	plugin_action_listCmd.Flags().String("plugin", "", "")
	plugin_actionCmd.AddCommand(plugin_action_listCmd)
}
