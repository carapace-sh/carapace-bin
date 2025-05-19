package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/terraform"
	"github.com/spf13/cobra"
)

var workspace_newCmd = &cobra.Command{
	Use:   "new [OPTIONS] NAME",
	Short: "Create a new workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_newCmd).Standalone()

	workspace_newCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	workspace_newCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	workspace_newCmd.Flags().StringS("state", "state", "", "Copy an existing state file into the new workspace.")
	workspaceCmd.AddCommand(workspace_newCmd)

	workspace_newCmd.Flag("lock-timeout").NoOptDefVal = " "
	workspace_newCmd.Flag("state").NoOptDefVal = " "

	carapace.Gen(workspace_newCmd).FlagCompletion(carapace.ActionMap{
		"state": carapace.ActionFiles(),
	})

	carapace.Gen(workspace_newCmd).PositionalCompletion(
		terraform.ActionWorkspaces(),
	)
}
