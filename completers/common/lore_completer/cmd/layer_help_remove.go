package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var layer_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a repository layer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(layer_help_removeCmd).Standalone()

	layer_helpCmd.AddCommand(layer_help_removeCmd)
}
