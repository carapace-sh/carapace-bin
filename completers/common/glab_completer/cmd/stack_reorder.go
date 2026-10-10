package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_reorderCmd = &cobra.Command{
	Use:   "reorder",
	Short: "Reorder a stack of diffs. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_reorderCmd).Standalone()

	stack_reorderCmd.Flags().Bool("abort", false, "Abort a reorder and restore original branch state.")
	stack_reorderCmd.Flags().Bool("continue", false, "Continue a reorder after resolving conflicts.")
	stackCmd.AddCommand(stack_reorderCmd)
}
