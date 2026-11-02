package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_executionModes_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the list of executionModes that would be run for a given configuration (default skaffold configuration, specific module, specific profile, etc).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_executionModes_listCmd).Standalone()

	applyCommonFlagCompletion(inspect_executionModes_listCmd)
	inspect_executionModesCmd.AddCommand(inspect_executionModes_listCmd)
}
