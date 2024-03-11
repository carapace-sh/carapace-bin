package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Short:   "List stacks",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_lsCmd).Standalone()

	stack_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	stackCmd.AddCommand(stack_lsCmd)
}
