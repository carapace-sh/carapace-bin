package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save your progress within a stacked diff. (EXPERIMENTAL.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_saveCmd).Standalone()

	stack_saveCmd.Flags().StringP("description", "d", "", "Description of the change.")
	stack_saveCmd.Flags().StringP("message", "m", "", "Alias for the description flag.")
	stackCmd.AddCommand(stack_saveCmd)
}
