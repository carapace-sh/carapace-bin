package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_layer_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List repository layers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_layer_listCmd).Standalone()

	help_layerCmd.AddCommand(help_layer_listCmd)
}
