package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var providers_schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Show schemas for the providers used in the configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providers_schemaCmd).Standalone()

	providers_schemaCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	providers_schemaCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	providers_schemaCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	providers_schemaCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	providers_schemaCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	providers_schemaCmd.Flags().BoolS("json", "json", false, "Prints out a json representation of the providers used in the current configuration.")
	providers_schemaCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	providers_schemaCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	providers_schemaCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")

	providersCmd.AddCommand(providers_schemaCmd)

	carapace.Gen(providers_schemaCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"var-file":    carapace.ActionFiles(),
	})
}
