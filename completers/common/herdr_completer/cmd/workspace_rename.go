package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var workspace_renameCmd = &cobra.Command{
	Use:   "rename <workspace_id> <label>",
	Short: "Rename a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_renameCmd).Standalone()

	workspaceCmd.AddCommand(workspace_renameCmd)

	carapace.Gen(workspace_renameCmd).PositionalCompletion(herdr.ActionWorkspaces())
}
