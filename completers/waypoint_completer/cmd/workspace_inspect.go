package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var workspace_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Output information for a given Workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_inspectCmd).Standalone()

	addGlobalOptions(workspace_inspectCmd)

	workspaceCmd.AddCommand(workspace_inspectCmd)
}
