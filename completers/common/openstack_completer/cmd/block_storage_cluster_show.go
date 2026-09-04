package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var block_storage_cluster_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show detailed information for a block storage cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(block_storage_cluster_showCmd).Standalone()

	block_storage_cluster_showCmd.Flags().String("binary", "", "Service binary.")
	block_storage_cluster_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	block_storage_cluster_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	block_storage_cluster_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	block_storage_cluster_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	block_storage_cluster_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	block_storage_cluster_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	block_storage_cluster_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	block_storage_cluster_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	block_storage_clusterCmd.AddCommand(block_storage_cluster_showCmd)
}
