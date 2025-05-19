package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_reorderCmd = &cobra.Command{
	Use:   "reorder",
	Short: "Reorder a stack of merge requests. (EXPERIMENTAL.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_reorderCmd).Standalone()

	stackCmd.AddCommand(stack_reorderCmd)
}
