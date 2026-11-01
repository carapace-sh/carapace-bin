package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tofu"
	"github.com/spf13/cobra"
)

var providers_mirrorCmd = &cobra.Command{
	Use:   "mirror [options] <target-dir>",
	Short: "Save local copies of all required provider plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providers_mirrorCmd).Standalone()

	providers_mirrorCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	providers_mirrorCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	providers_mirrorCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	providers_mirrorCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	providers_mirrorCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	providers_mirrorCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	providers_mirrorCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	providers_mirrorCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	providers_mirrorCmd.Flags().StringArrayS("platform", "platform", nil, "Choose which target platform to build a mirror for.")
	providers_mirrorCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	providers_mirrorCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	providersCmd.AddCommand(providers_mirrorCmd)

	carapace.Gen(providers_mirrorCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"json-into":   carapace.ActionFiles(),
		"platform":    tofu.ActionPlatforms(),
		"var-file":    carapace.ActionFiles(),
	})

	carapace.Gen(providers_mirrorCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
