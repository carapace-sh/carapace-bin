package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var workspace_newCmd = &cobra.Command{
	Use:   "new [options] <NAME>",
	Short: "Create a new workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_newCmd).Standalone()

	workspace_newCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	workspace_newCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	workspace_newCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	workspace_newCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	workspace_newCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	workspace_newCmd.Flags().BoolS("json", "json", false, "The output of the command is printed in json format.")
	workspace_newCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	workspace_newCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	workspace_newCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	workspace_newCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	workspace_newCmd.Flags().StringS("state", "state", "", "Copy an existing state file into the new workspace.")
	workspace_newCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	workspace_newCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	workspaceCmd.AddCommand(workspace_newCmd)

	carapace.Gen(workspace_newCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"json-into":   carapace.ActionFiles(),
		"state":       carapace.ActionFiles(),
		"var-file":    carapace.ActionFiles(),
	})
}
