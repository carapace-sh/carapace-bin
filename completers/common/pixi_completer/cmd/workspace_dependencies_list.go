package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_dependencies_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List dependencies in the workspace dependencies table",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_dependencies_listCmd).Standalone()

	workspace_dependenciesCmd.AddCommand(workspace_dependencies_listCmd)
}
