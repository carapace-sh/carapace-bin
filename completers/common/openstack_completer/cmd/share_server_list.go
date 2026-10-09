package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_server_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all share servers (Admin only).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_server_listCmd).Standalone()

	share_server_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_server_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_server_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_server_listCmd.Flags().String("host", "", "Filter results by name of host.")
	share_server_listCmd.Flags().String("identifier", "", "Identifier of the share server in the share back end.")
	share_server_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_server_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_server_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_server_listCmd.Flags().String("project", "", "Filter results by project name or ID.")
	share_server_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	share_server_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_server_listCmd.Flags().String("share-network", "", "Filter results by share network name or ID.")
	share_server_listCmd.Flags().String("share-network-subnet", "", "Filter results by share network subnet that the share server's network allocation exists within.")
	share_server_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_server_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_server_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_server_listCmd.Flags().String("source-share-server-id", "", "Share server ID to be used as a filter.")
	share_server_listCmd.Flags().String("status", "", "Filter results by status.")
	share_serverCmd.AddCommand(share_server_listCmd)
}
