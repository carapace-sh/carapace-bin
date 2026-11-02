package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_executionModesCmd = &cobra.Command{
	Use:   "executionModes",
	Short: "View skaffold executionMode information defined in the specified skaffold configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_executionModesCmd).Standalone()

	applyCommonFlagCompletion(inspect_executionModesCmd)
	inspectCmd.AddCommand(inspect_executionModesCmd)
}
