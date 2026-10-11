package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pixi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var workspace_dependencies_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove dependencies from the workspace dependencies table",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_dependencies_removeCmd).Standalone()

	workspace_dependenciesCmd.AddCommand(workspace_dependencies_removeCmd)

	carapace.Gen(workspace_dependencies_removeCmd).PositionalAnyCompletion(
		action.ActionWorkspaceDependencies(workspace_dependencies_removeCmd),
	)
}
