package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [options] PATH",
	Short: "Install or upgrade remote OpenTofu modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	getCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	getCmd.Flags().BoolS("no-color", "no-color", false, "Disable text coloring in the output.")
	getCmd.Flags().StringS("test-directory", "test-directory", "", "Set the OpenTofu test directory, defaults to \"tests\".")
	getCmd.Flags().BoolS("update", "update", false, "Check already-downloaded modules for available updates.")
	getCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	getCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(getCmd)

	getCmd.Flag("json-into").NoOptDefVal = " "
	getCmd.Flag("test-directory").NoOptDefVal = " "
	getCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"json-into":      carapace.ActionFiles(),
		"test-directory": carapace.ActionDirectories(),
		"var-file":       carapace.ActionFiles(),
	})

	carapace.Gen(getCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
