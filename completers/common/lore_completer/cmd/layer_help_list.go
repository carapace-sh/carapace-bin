package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var layer_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List repository layers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(layer_help_listCmd).Standalone()

	layer_helpCmd.AddCommand(layer_help_listCmd)
}
