package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/terraform"
	"github.com/spf13/cobra"
)

var workspace_deleteCmd = &cobra.Command{
	Use:   "delete [OPTIONS] NAME",
	Short: "Delete a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_deleteCmd).Standalone()

	workspace_deleteCmd.Flags().BoolS("force", "force", false, "Remove even a non-empty workspace.")
	workspace_deleteCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	workspace_deleteCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	workspaceCmd.AddCommand(workspace_deleteCmd)

	workspace_deleteCmd.Flag("lock-timeout").NoOptDefVal = " "

	carapace.Gen(workspace_deleteCmd).PositionalCompletion(
		terraform.ActionWorkspaces(),
	)
}
