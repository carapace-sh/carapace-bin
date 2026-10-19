package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tofu"
	"github.com/spf13/cobra"
)

var providers_lockCmd = &cobra.Command{
	Use:   "lock [options] [providers...]",
	Short: "Write out dependency locks for the configured providers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providers_lockCmd).Standalone()

	providers_lockCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	providers_lockCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	providers_lockCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	providers_lockCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	providers_lockCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	providers_lockCmd.Flags().StringS("fs-mirror", "fs-mirror", "", "Consult the given filesystem mirror directory instead of the configured provider installations.")
	providers_lockCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	providers_lockCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	providers_lockCmd.Flags().StringS("net-mirror", "net-mirror", "", "Consult the given network mirror (given as a base URL) instead of the configured provider installations.")
	providers_lockCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	providers_lockCmd.Flags().StringS("oci-mirror", "oci-mirror", "", "Consult the given OCI registry mirror (given as a template) instead of the configured provider installations.")
	providers_lockCmd.Flags().StringArrayS("platform", "platform", nil, "Choose a target platform to request package checksums for.")
	providers_lockCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	providers_lockCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	providersCmd.AddCommand(providers_lockCmd)

	carapace.Gen(providers_lockCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"fs-mirror":   carapace.ActionDirectories(),
		"json-into":   carapace.ActionFiles(),
		"platform":    tofu.ActionPlatforms(),
		"var-file":    carapace.ActionFiles(),
	})

	carapace.Gen(providers_lockCmd).PositionalAnyCompletion(
		tofu.ActionProviders(),
	)
}
