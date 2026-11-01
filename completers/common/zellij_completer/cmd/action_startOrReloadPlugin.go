package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_startOrReloadPluginCmd = &cobra.Command{
	Use:   "start-or-reload-plugin",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_startOrReloadPluginCmd).Standalone()

	action_startOrReloadPluginCmd.Flags().StringP("configuration", "c", "", "")
	action_startOrReloadPluginCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_startOrReloadPluginCmd)
}
