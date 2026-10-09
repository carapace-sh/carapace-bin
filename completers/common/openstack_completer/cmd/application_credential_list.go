package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var application_credential_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List application credentials",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(application_credential_listCmd).Standalone()

	application_credential_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	application_credential_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	application_credential_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	application_credential_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	application_credential_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	application_credential_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	application_credential_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	application_credential_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	application_credential_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	application_credential_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	application_credential_listCmd.Flags().String("user", "", "User whose application credentials to list (name or ID)")
	application_credential_listCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	application_credentialCmd.AddCommand(application_credential_listCmd)
}
