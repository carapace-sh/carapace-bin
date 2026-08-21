package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_security_service_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List security services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_security_service_listCmd).Standalone()

	share_security_service_listCmd.Flags().Bool("all-projects", false, "Display information from all projects (Admin only).")
	share_security_service_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_security_service_listCmd.Flags().String("default-ad-site", "", "Filter results by security service default_ad_site.")
	share_security_service_listCmd.Flags().Bool("detail", false, "Show detailed information about filtered security services.")
	share_security_service_listCmd.Flags().String("dns-ip", "", "Filter results by DNS IP address used inside project's network.")
	share_security_service_listCmd.Flags().String("domain", "", "Filter results by security service domain.")
	share_security_service_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_security_service_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_security_service_listCmd.Flags().String("limit", "", "Limit the number of security services returned")
	share_security_service_listCmd.Flags().String("marker", "", "The last security service ID of the previous page")
	share_security_service_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_security_service_listCmd.Flags().String("name", "", "Filter results by security service name.")
	share_security_service_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_security_service_listCmd.Flags().String("ou", "", "Filter results by security service OU (Organizational Unit).")
	share_security_service_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_security_service_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_security_service_listCmd.Flags().String("server", "", "Filter results by security service IP address or hostname.")
	share_security_service_listCmd.Flags().String("share-network", "", "Filter results by share network name or ID.")
	share_security_service_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_security_service_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_security_service_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_security_service_listCmd.Flags().String("status", "", "Filter results by status.")
	share_security_service_listCmd.Flags().String("type", "", "Filter results by security service type.")
	share_security_service_listCmd.Flags().String("user", "", "Filter results by security service user or group used by project.")
	share_security_serviceCmd.AddCommand(share_security_service_listCmd)
}
