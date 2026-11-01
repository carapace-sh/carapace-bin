package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var layer_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List repository layers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(layer_listCmd).Standalone()

	layer_listCmd.Flags().BoolP("help", "h", false, "Print help")
	layerCmd.AddCommand(layer_listCmd)
}
