package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_configDependencies_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add config dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_configDependencies_addCmd).Standalone()

	inspect_configDependencies_addCmd.Flags().StringSliceP("module", "m", nil, "Names of modules to filter target action by.")
	applyCommonFlagCompletion(inspect_configDependencies_addCmd)
	inspect_configDependenciesCmd.AddCommand(inspect_configDependencies_addCmd)
}
