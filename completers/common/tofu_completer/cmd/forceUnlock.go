package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var forceUnlockCmd = &cobra.Command{
	Use:   "force-unlock [options] <LOCK_ID>",
	Short: "Release a stuck lock on the current workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forceUnlockCmd).Standalone()

	forceUnlockCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	forceUnlockCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	forceUnlockCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	forceUnlockCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	forceUnlockCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	forceUnlockCmd.Flags().BoolS("force", "force", false, "Don't ask for input for unlock confirmation.")
	forceUnlockCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	forceUnlockCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	forceUnlockCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	forceUnlockCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	forceUnlockCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(forceUnlockCmd)

	carapace.Gen(forceUnlockCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"json-into":   carapace.ActionFiles(),
		"var-file":    carapace.ActionFiles(),
	})
}
