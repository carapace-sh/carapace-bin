package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tofu"
	"github.com/spf13/cobra"
)

var providers_lockCmd = &cobra.Command{
	Use:   "lock [options]",
	Short: "Write out dependency locks for the configured providers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providers_lockCmd).Standalone()

	providers_lockCmd.Flags().StringS("fs-mirror", "fs-mirror", "", "Consult the given filesystem mirror directory instead of the configured provider installations.")
	providers_lockCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	providers_lockCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	providers_lockCmd.Flags().StringS("net-mirror", "net-mirror", "", "Consult the given network mirror (given as a base URL) instead of the configured provider installations.")
	providers_lockCmd.Flags().StringS("platform", "platform", "", "Choose a target platform to request package checksums for.")
	providers_lockCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	providers_lockCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	providersCmd.AddCommand(providers_lockCmd)

	providers_lockCmd.Flag("fs-mirror").NoOptDefVal = " "
	providers_lockCmd.Flag("json-into").NoOptDefVal = " "
	providers_lockCmd.Flag("net-mirror").NoOptDefVal = " "
	providers_lockCmd.Flag("platform").NoOptDefVal = " "
	providers_lockCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(providers_lockCmd).FlagCompletion(carapace.ActionMap{
		"fs-mirror": carapace.ActionDirectories(),
		"json-into": carapace.ActionFiles(),
		"platform":  tofu.ActionPlatforms(),
		"var-file":  carapace.ActionFiles(),
	})
}
