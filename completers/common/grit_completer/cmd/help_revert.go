package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Undo commits with new inverse commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revertCmd).Standalone()

	helpCmd.AddCommand(help_revertCmd)
}
