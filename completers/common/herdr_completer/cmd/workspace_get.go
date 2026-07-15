package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var workspace_getCmd = &cobra.Command{
	Use:   "get <workspace_id>",
	Short: "Show a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_getCmd).Standalone()

	workspaceCmd.AddCommand(workspace_getCmd)

	carapace.Gen(workspace_getCmd).PositionalCompletion(action.ActionWorkspaces())
}
