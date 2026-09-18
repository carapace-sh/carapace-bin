package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var application_credential_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display application credential details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(application_credential_showCmd).Standalone()

	application_credential_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	application_credential_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	application_credential_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	application_credential_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	application_credential_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	application_credential_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	application_credential_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	application_credential_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	application_credentialCmd.AddCommand(application_credential_showCmd)
}
