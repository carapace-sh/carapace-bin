package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var credential_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List credentials",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(credential_listCmd).Standalone()

	credential_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	credential_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	credential_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	credential_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	credential_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	credential_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	credential_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	credential_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	credential_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	credential_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	credential_listCmd.Flags().String("type", "", "Filter credentials by type: cert, ec2, totp and so on")
	credential_listCmd.Flags().String("user", "", "Filter credentials by <user> (name or ID)")
	credential_listCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	credentialCmd.AddCommand(credential_listCmd)
}
