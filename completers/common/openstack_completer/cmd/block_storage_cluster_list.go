package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var block_storage_cluster_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List block storage clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(block_storage_cluster_listCmd).Standalone()

	block_storage_cluster_listCmd.Flags().String("binary", "", "Cluster binary.")
	block_storage_cluster_listCmd.Flags().String("cluster", "", "Filter by cluster name, without backend will list all clustered services from the same cluster.")
	block_storage_cluster_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	block_storage_cluster_listCmd.Flags().Bool("disabled", false, "Filter by disabled status.")
	block_storage_cluster_listCmd.Flags().Bool("down", false, "Filter by down status.")
	block_storage_cluster_listCmd.Flags().Bool("enabled", false, "Filter by enabled status.")
	block_storage_cluster_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	block_storage_cluster_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	block_storage_cluster_listCmd.Flags().Bool("long", false, "List additional fields in output")
	block_storage_cluster_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	block_storage_cluster_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	block_storage_cluster_listCmd.Flags().String("num-down-hosts", "", "Filter by number of hosts that are down.")
	block_storage_cluster_listCmd.Flags().String("num-hosts", "", "Filter by number of hosts in the cluster.")
	block_storage_cluster_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	block_storage_cluster_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	block_storage_cluster_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	block_storage_cluster_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	block_storage_cluster_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	block_storage_cluster_listCmd.Flags().Bool("up", false, "Filter by up status.")
	block_storage_clusterCmd.AddCommand(block_storage_cluster_listCmd)
}
