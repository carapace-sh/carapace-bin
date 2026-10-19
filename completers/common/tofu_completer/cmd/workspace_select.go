package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tofu"
	"github.com/spf13/cobra"
)

var workspace_selectCmd = &cobra.Command{
	Use:   "select [options] <NAME>",
	Short: "Select a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_selectCmd).Standalone()

	workspace_selectCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	workspace_selectCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	workspace_selectCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	workspace_selectCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	workspace_selectCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	workspace_selectCmd.Flags().BoolS("json", "json", false, "The output of the command is printed in json format.")
	workspace_selectCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	workspace_selectCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	workspace_selectCmd.Flags().BoolS("or-create", "or-create", false, "Create the OpenTofu workspace if it doesn't exist.")
	workspace_selectCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	workspace_selectCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	workspaceCmd.AddCommand(workspace_selectCmd)

	carapace.Gen(workspace_selectCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"json-into":   carapace.ActionFiles(),
		"var-file":    carapace.ActionFiles(),
	})

	carapace.Gen(workspace_selectCmd).PositionalCompletion(
		tofu.ActionWorkspaces(),
	)
}
