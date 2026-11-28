package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_testsCmd = &cobra.Command{
	Use:   "tests",
	Short: "View skaffold test information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_testsCmd).Standalone()

	applyCommonFlagCompletion(inspect_testsCmd)
	inspectCmd.AddCommand(inspect_testsCmd)
}
