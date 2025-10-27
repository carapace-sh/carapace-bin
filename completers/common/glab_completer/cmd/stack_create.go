package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new stacked diff. (EXPERIMENTAL)",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_createCmd).Standalone()

	stackCmd.AddCommand(stack_createCmd)
}
