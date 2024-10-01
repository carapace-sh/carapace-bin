package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Moves to any selected entry in the stack. (EXPERIMENTAL.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_moveCmd).Standalone()

	stackCmd.AddCommand(stack_moveCmd)
}
