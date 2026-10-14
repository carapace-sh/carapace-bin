package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var credential_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display credential details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(credential_showCmd).Standalone()

	credential_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	credential_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	credential_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	credential_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	credential_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	credential_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	credential_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	credential_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	credentialCmd.AddCommand(credential_showCmd)
}
