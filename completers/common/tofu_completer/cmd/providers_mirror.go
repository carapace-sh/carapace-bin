package cmd

import (
	"github.com/carapace-sh/carapace"
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

	providers_mirrorCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	providers_mirrorCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	providers_mirrorCmd.Flags().StringS("platform", "platform", "", "Choose which target platform to build a mirror for.")
	providers_mirrorCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	providers_mirrorCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	providersCmd.AddCommand(providers_mirrorCmd)

	providers_mirrorCmd.Flag("json-into").NoOptDefVal = " "
	providers_mirrorCmd.Flag("platform").NoOptDefVal = " "
	providers_mirrorCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(providers_mirrorCmd).FlagCompletion(carapace.ActionMap{
		"json-into": carapace.ActionFiles(),
		"platform":  tofu.ActionPlatforms(),
		"var-file":  carapace.ActionFiles(),
	})

	carapace.Gen(providers_mirrorCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
