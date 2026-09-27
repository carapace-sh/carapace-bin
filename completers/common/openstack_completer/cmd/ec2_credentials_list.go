package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_credentials_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List EC2 credentials",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_credentials_listCmd).Standalone()

	ec2_credentials_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	ec2_credentials_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	ec2_credentials_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	ec2_credentials_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	ec2_credentials_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	ec2_credentials_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	ec2_credentials_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	ec2_credentials_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	ec2_credentials_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	ec2_credentials_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	ec2_credentials_listCmd.Flags().String("user", "", "Filter list by user (name or ID)")
	ec2_credentials_listCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	ec2_credentialsCmd.AddCommand(ec2_credentials_listCmd)
}
