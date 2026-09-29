package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var outputCmd = &cobra.Command{
	Use:   "output [options] [NAME]",
	Short: "Show output values from your root module",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outputCmd).Standalone()

	outputCmd.Flags().BoolS("json", "json", false, "If specified, machine readable output will be printed in JSON format.")
	outputCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	outputCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	outputCmd.Flags().BoolS("raw", "raw", false, "For value types that can be automatically converted to a string, will print the raw string directly.")
	outputCmd.Flags().BoolS("show-sensitive", "show-sensitive", false, "If specified, sensitive values will be displayed.")
	outputCmd.Flags().StringS("state", "state", "", "Path to the state file to read. Defaults to \"terraform.tfstate\".")
	outputCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	outputCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(outputCmd)

	outputCmd.Flag("json-into").NoOptDefVal = " "
	outputCmd.Flag("state").NoOptDefVal = " "
	outputCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(outputCmd).FlagCompletion(carapace.ActionMap{
		"json-into": carapace.ActionFiles(),
		"state":     carapace.ActionFiles(),
		"var-file":  carapace.ActionFiles(),
	})

	carapace.Gen(outputCmd).PositionalCompletion(
		action.ActionOutputs(outputCmd),
	)
}
