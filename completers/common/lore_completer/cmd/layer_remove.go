package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var layer_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a repository layer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(layer_removeCmd).Standalone()

	layer_removeCmd.Flags().BoolP("help", "h", false, "Print help")
	layer_removeCmd.Flags().Bool("purge", false, "Also delete untracked files and all directories inside the layer mount")
	layerCmd.AddCommand(layer_removeCmd)
}
