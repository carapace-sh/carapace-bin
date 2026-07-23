package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var workspace_closeCmd = &cobra.Command{
	Use:   "close <workspace_id>",
	Short: "Close a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_closeCmd).Standalone()

	workspaceCmd.AddCommand(workspace_closeCmd)

	carapace.Gen(workspace_closeCmd).PositionalCompletion(herdr.ActionWorkspaces())
}
