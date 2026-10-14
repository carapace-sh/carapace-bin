package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [options] [path]",
	Short: "Show the current state or a saved plan",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().BoolS("config", "config", false, "Show the current configuration (requires -json).")
	showCmd.Flags().BoolS("json", "json", false, "Show the information in a machine-readable form.")
	showCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	showCmd.Flags().BoolS("no-color", "no-color", false, "Disable terminal escape sequences.")
	showCmd.Flags().StringS("plan", "plan", "", "The plan from a saved plan file.")
	showCmd.Flags().BoolS("show-sensitive", "show-sensitive", false, "If specified, sensitive values will be displayed.")
	showCmd.Flags().BoolS("state", "state", false, "The latest state snapshot, if any.")
	showCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	showCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(showCmd)

	showCmd.Flag("json-into").NoOptDefVal = " "
	showCmd.Flag("plan").NoOptDefVal = " "
	showCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(showCmd).FlagCompletion(carapace.ActionMap{
		"json-into": carapace.ActionFiles(),
		"plan":      carapace.ActionFiles(),
		"var-file":  carapace.ActionFiles(),
	})

	carapace.Gen(showCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
