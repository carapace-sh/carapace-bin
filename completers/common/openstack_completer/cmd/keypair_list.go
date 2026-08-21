package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keypair_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List key fingerprints",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keypair_listCmd).Standalone()

	keypair_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	keypair_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	keypair_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	keypair_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	keypair_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	keypair_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	keypair_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	keypair_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	keypair_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	keypair_listCmd.Flags().String("project", "", "Show keypairs for all users associated with project (admin only) (name or ID).")
	keypair_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	keypair_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	keypair_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	keypair_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	keypair_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	keypair_listCmd.Flags().String("user", "", "Show keypairs for another user (admin only) (name or ID).")
	keypair_listCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	keypairCmd.AddCommand(keypair_listCmd)
}
