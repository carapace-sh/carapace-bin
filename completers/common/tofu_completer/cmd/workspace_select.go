package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tofu"
	"github.com/spf13/cobra"
)

var workspace_selectCmd = &cobra.Command{
	Use:   "select NAME",
	Short: "Select a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_selectCmd).Standalone()

	workspace_selectCmd.Flags().BoolS("json", "json", false, "The output of the command is printed in json format.")
	workspace_selectCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	workspace_selectCmd.Flags().BoolS("or-create", "or-create", false, "Create the OpenTofu workspace if it doesn't exist.")
	workspace_selectCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	workspace_selectCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	workspaceCmd.AddCommand(workspace_selectCmd)

	workspace_selectCmd.Flag("json-into").NoOptDefVal = " "
	workspace_selectCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(workspace_selectCmd).FlagCompletion(carapace.ActionMap{
		"json-into": carapace.ActionFiles(),
		"var-file":  carapace.ActionFiles(),
	})

	carapace.Gen(workspace_selectCmd).PositionalCompletion(
		tofu.ActionWorkspaces(),
	)
}
