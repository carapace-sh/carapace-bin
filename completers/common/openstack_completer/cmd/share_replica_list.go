package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_replica_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List share replicas",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_replica_listCmd).Standalone()

	share_replica_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_replica_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_replica_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_replica_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_replica_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_replica_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_replica_listCmd.Flags().String("property", "", "Filter replicas having a given metadata key=value property.")
	share_replica_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_replica_listCmd.Flags().String("share", "", "Name or ID of the share to list replicas for.")
	share_replica_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_replica_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_replica_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_replicaCmd.AddCommand(share_replica_listCmd)
}
