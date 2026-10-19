package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [options] [arguments...]",
	Short: "Show the current state or a saved plan",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	showCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	showCmd.Flags().BoolS("config", "config", false, "Show the current configuration (requires -json).")
	showCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	showCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	showCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	showCmd.Flags().BoolS("json", "json", false, "Show the information in a machine-readable form.")
	showCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	showCmd.Flags().StringS("module", "module", "", "Show the specified module configuration (requires -json).")
	showCmd.Flags().BoolS("no-color", "no-color", false, "Disable terminal escape sequences.")
	showCmd.Flags().StringS("plan", "plan", "", "The plan from a saved plan file.")
	showCmd.Flags().BoolS("show-sensitive", "show-sensitive", false, "If specified, sensitive values will be displayed.")
	showCmd.Flags().BoolS("state", "state", false, "The latest state snapshot, if any.")
	showCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	showCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"json-into":   carapace.ActionFiles(),
		"module":      carapace.ActionDirectories(),
		"plan":        carapace.ActionFiles(),
		"var-file":    carapace.ActionFiles(),
	})

	carapace.Gen(showCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
