package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_prevCmd = &cobra.Command{
	Use:   "prev",
	Short: "Moves to the previous diff in the stack. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_prevCmd).Standalone()

	stackCmd.AddCommand(stack_prevCmd)
}
