package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the global manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_editCmd).Standalone()

	globalCmd.AddCommand(global_editCmd)
}
