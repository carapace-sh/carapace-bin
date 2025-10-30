package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo the last operation by reverting to the previous snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_undoCmd).Standalone()

	helpCmd.AddCommand(help_undoCmd)
}
