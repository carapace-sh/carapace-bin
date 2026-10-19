package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version [options]",
	Short: "Show the current OpenTofu version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	versionCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	versionCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	versionCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	versionCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	versionCmd.Flags().BoolS("json", "json", false, "Output the version information as a JSON object.")
	versionCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	rootCmd.AddCommand(versionCmd)

	carapace.Gen(versionCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
	})
}
