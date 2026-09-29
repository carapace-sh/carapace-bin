package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the name of the current workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_showCmd).Standalone()

	workspace_showCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	workspace_showCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	workspaceCmd.AddCommand(workspace_showCmd)

	workspace_showCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(workspace_showCmd).FlagCompletion(carapace.ActionMap{
		"var-file": carapace.ActionFiles(),
	})
}
