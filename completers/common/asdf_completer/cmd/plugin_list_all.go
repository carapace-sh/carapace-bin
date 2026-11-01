package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_list_allCmd = &cobra.Command{
	Use:   "all",
	Short: "List plugins registered on asdf-plugins repository with URLs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_list_allCmd).Standalone()

	plugin_listCmd.AddCommand(plugin_list_allCmd)
}
