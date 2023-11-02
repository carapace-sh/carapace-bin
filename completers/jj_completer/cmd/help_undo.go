package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo an operation (shortcut for `jj op undo`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_undoCmd).Standalone()

	helpCmd.AddCommand(help_undoCmd)
}
