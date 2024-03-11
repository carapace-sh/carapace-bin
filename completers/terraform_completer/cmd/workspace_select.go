package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/terraform"
	"github.com/spf13/cobra"
)

var workspace_selectCmd = &cobra.Command{
	Use:   "select NAME",
	Short: "Select a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_selectCmd).Standalone()

	workspace_selectCmd.Flags().BoolS("or-create", "or-create", false, "Create the Terraform workspace if it doesn't exist")
	workspaceCmd.AddCommand(workspace_selectCmd)

	carapace.Gen(workspace_selectCmd).PositionalCompletion(
		terraform.ActionWorkspaces(),
	)
}
