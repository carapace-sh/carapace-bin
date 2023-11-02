package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var operation_help_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Create a new operation that restores the repo to an earlier state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operation_help_restoreCmd).Standalone()

	operation_helpCmd.AddCommand(operation_help_restoreCmd)
}
