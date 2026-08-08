package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_inferCmd = &cobra.Command{
	Use:   "infer <revision-range>",
	Short: "Add layers to a stack based on a range of commits. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_inferCmd).Standalone()

	stack_inferCmd.Flags().StringP("name", "n", "", "Name for the new stack (used when creating a stack)")
	stackCmd.AddCommand(stack_inferCmd)
}
