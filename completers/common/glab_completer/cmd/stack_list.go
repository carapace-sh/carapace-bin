package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all diffs in the stack. (EXPERIMENTAL)",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_listCmd).Standalone()

	stackCmd.AddCommand(stack_listCmd)
}
