package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_firstCmd = &cobra.Command{
	Use:   "first",
	Short: "Moves to the first diff in the stack. (EXPERIMENTAL.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_firstCmd).Standalone()

	stackCmd.AddCommand(stack_firstCmd)
}
