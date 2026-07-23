package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var workspace_focusCmd = &cobra.Command{
	Use:   "focus <workspace_id>",
	Short: "Focus a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_focusCmd).Standalone()

	workspaceCmd.AddCommand(workspace_focusCmd)

	carapace.Gen(workspace_focusCmd).PositionalCompletion(herdr.ActionWorkspaces())
}
