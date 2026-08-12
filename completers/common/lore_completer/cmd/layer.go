package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var layerCmd = &cobra.Command{
	Use:   "layer",
	Short: "Layer commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(layerCmd).Standalone()

	layerCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(layerCmd)
}
