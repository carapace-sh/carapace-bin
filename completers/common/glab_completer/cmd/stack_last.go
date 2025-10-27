package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_lastCmd = &cobra.Command{
	Use:   "last",
	Short: "Moves to the last diff in the stack. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_lastCmd).Standalone()

	stackCmd.AddCommand(stack_lastCmd)
}
