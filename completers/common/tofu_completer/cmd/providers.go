package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var providersCmd = &cobra.Command{
	Use:   "providers",
	Short: "Show the providers required for this configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providersCmd).Standalone()

	providersCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	providersCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	providersCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	providersCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	providersCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	providersCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	providersCmd.Flags().StringS("test-directory", "test-directory", "", "Set the OpenTofu test directory, defaults to \"tests\".")
	providersCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	providersCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(providersCmd)

	carapace.Gen(providersCmd).FlagCompletion(carapace.ActionMap{
		"deprecation":    action.ActionDeprecationLevels(),
		"test-directory": carapace.ActionDirectories(),
		"var-file":       carapace.ActionFiles(),
	})
}
