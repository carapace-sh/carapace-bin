package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listDependenciesCmd = &cobra.Command{
	Use:     "list-dependencies",
	Short:   "Recursively show units which are required or wanted",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listDependenciesCmd).Standalone()

	rootCmd.AddCommand(listDependenciesCmd)

	carapace.Gen(listDependenciesCmd).PositionalAnyCompletion(
		action.ActionUnits(listDependenciesCmd).FilterArgs(),
	)
}
