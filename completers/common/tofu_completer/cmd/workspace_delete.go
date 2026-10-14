package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tofu"
	"github.com/spf13/cobra"
)

var workspace_deleteCmd = &cobra.Command{
	Use:   "delete [OPTIONS] NAME",
	Short: "Delete a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_deleteCmd).Standalone()

	workspace_deleteCmd.Flags().BoolS("force", "force", false, "Remove a workspace even if it is managing resources.")
	workspace_deleteCmd.Flags().BoolS("json", "json", false, "The output of the command is printed in json format.")
	workspace_deleteCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	workspace_deleteCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	workspace_deleteCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	workspace_deleteCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	workspace_deleteCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	workspaceCmd.AddCommand(workspace_deleteCmd)

	workspace_deleteCmd.Flag("json-into").NoOptDefVal = " "
	workspace_deleteCmd.Flag("lock-timeout").NoOptDefVal = " "
	workspace_deleteCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(workspace_deleteCmd).FlagCompletion(carapace.ActionMap{
		"json-into": carapace.ActionFiles(),
		"var-file":  carapace.ActionFiles(),
	})

	carapace.Gen(workspace_deleteCmd).PositionalCompletion(
		tofu.ActionWorkspaces(),
	)
}
