package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Moves to the next diff in the stack. (EXPERIMENTAL.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_nextCmd).Standalone()

	stackCmd.AddCommand(stack_nextCmd)
}
