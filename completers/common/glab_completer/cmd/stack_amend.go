package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_amendCmd = &cobra.Command{
	Use:   "amend",
	Short: "Save more changes to a stacked diff. (EXPERIMENTAL.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_amendCmd).Standalone()

	stack_amendCmd.Flags().StringP("description", "d", "", "a description of the change")
	stack_amendCmd.Flags().StringP("message", "m", "", "alias for the description flag")
	stackCmd.AddCommand(stack_amendCmd)
}
