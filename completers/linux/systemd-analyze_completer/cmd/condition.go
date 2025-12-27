package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var conditionCmd = &cobra.Command{
	Use:   "condition",
	Short: "Evaluate conditions and asserts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(conditionCmd).Standalone()

	rootCmd.AddCommand(conditionCmd)
}
