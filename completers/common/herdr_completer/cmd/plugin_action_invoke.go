package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_action_invokeCmd = &cobra.Command{
	Use:   "invoke <action_id>",
	Short: "Invoke a plugin action",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_action_invokeCmd).Standalone()

	plugin_action_invokeCmd.Flags().String("plugin", "", "")
	plugin_actionCmd.AddCommand(plugin_action_invokeCmd)
}
