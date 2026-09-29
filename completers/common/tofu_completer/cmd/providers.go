package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var providersCmd = &cobra.Command{
	Use:   "providers",
	Short: "Show the providers required for this configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providersCmd).Standalone()

	providersCmd.Flags().StringS("test-directory", "test-directory", "", "Set the OpenTofu test directory, defaults to \"tests\".")
	providersCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	providersCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(providersCmd)

	providersCmd.Flag("test-directory").NoOptDefVal = " "
	providersCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(providersCmd).FlagCompletion(carapace.ActionMap{
		"test-directory": carapace.ActionDirectories(),
		"var-file":       carapace.ActionFiles(),
	})
}
