package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var layer_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a repository layer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(layer_help_addCmd).Standalone()

	layer_helpCmd.AddCommand(layer_help_addCmd)
}
