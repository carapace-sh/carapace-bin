package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [options]",
	Short: "Install or upgrade remote OpenTofu modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	getCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	getCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	getCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	getCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	getCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	getCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	getCmd.Flags().BoolS("no-color", "no-color", false, "Disable text coloring in the output.")
	getCmd.Flags().StringS("test-directory", "test-directory", "", "Set the OpenTofu test directory, defaults to \"tests\".")
	getCmd.Flags().BoolS("update", "update", false, "Check already-downloaded modules for available updates.")
	getCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	getCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"deprecation":    action.ActionDeprecationLevels(),
		"json-into":      carapace.ActionFiles(),
		"test-directory": carapace.ActionDirectories(),
		"var-file":       carapace.ActionFiles(),
	})
}
