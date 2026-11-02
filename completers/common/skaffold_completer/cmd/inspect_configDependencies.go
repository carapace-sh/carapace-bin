package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_configDependenciesCmd = &cobra.Command{
	Use:   "config-dependencies",
	Short: "Interact with skaffold config dependency definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_configDependenciesCmd).Standalone()

	applyCommonFlagCompletion(inspect_configDependenciesCmd)
	inspectCmd.AddCommand(inspect_configDependenciesCmd)
}
