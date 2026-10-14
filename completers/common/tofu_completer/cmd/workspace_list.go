package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_listCmd).Standalone()

	workspace_listCmd.Flags().BoolS("json", "json", false, "The output of the command is printed in json format.")
	workspace_listCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	workspace_listCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	workspace_listCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	workspaceCmd.AddCommand(workspace_listCmd)

	workspace_listCmd.Flag("json-into").NoOptDefVal = " "
	workspace_listCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(workspace_listCmd).FlagCompletion(carapace.ActionMap{
		"json-into": carapace.ActionFiles(),
		"var-file":  carapace.ActionFiles(),
	})
}
