package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_layer_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a repository layer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_layer_removeCmd).Standalone()

	help_layerCmd.AddCommand(help_layer_removeCmd)
}
