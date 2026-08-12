package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_layerCmd = &cobra.Command{
	Use:   "layer",
	Short: "Layer commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_layerCmd).Standalone()

	helpCmd.AddCommand(help_layerCmd)
}
