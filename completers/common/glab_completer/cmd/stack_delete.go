package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_deleteCmd = &cobra.Command{
	Use:   "delete [<stack-name>]",
	Short: "Delete a stack. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_deleteCmd).Standalone()

	stack_deleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt.")
	stackCmd.AddCommand(stack_deleteCmd)
}
