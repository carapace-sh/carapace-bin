package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_log_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List plugin command logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_log_listCmd).Standalone()

	plugin_log_listCmd.Flags().String("limit", "", "")
	plugin_log_listCmd.Flags().String("plugin", "", "")
	plugin_logCmd.AddCommand(plugin_log_listCmd)
}
