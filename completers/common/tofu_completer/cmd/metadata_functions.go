package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var metadata_functionsCmd = &cobra.Command{
	Use:   "functions [options]",
	Short: "Show signatures and descriptions for the available functions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(metadata_functionsCmd).Standalone()

	metadata_functionsCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	metadata_functionsCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	metadata_functionsCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	metadata_functionsCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	metadata_functionsCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	metadata_functionsCmd.Flags().BoolS("json", "json", false, "Prints out a json representation of the available function signatures")
	metadata_functionsCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	metadataCmd.AddCommand(metadata_functionsCmd)

	carapace.Gen(metadata_functionsCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
	})
}
