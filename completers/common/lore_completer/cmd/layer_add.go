package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var layer_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a repository layer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(layer_addCmd).Standalone()

	layer_addCmd.Flags().BoolP("help", "h", false, "Print help")
	layer_addCmd.Flags().String("metadata", "", "Metadata key to use for matching revisions")
	layerCmd.AddCommand(layer_addCmd)
}
