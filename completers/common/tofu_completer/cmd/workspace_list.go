package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var workspace_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_listCmd).Standalone()

	workspace_listCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	workspace_listCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	workspace_listCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	workspace_listCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	workspace_listCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	workspace_listCmd.Flags().BoolS("json", "json", false, "The output of the command is printed in json format.")
	workspace_listCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	workspace_listCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	workspace_listCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	workspace_listCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	workspaceCmd.AddCommand(workspace_listCmd)

	carapace.Gen(workspace_listCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"json-into":   carapace.ActionFiles(),
		"var-file":    carapace.ActionFiles(),
	})
}
