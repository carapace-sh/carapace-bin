package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_server_adoptCmd = &cobra.Command{
	Use:   "adopt",
	Short: "Adopt share server not handled by Manila (Admin only).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_server_adoptCmd).Standalone()

	share_server_adoptCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_server_adoptCmd.Flags().String("driver-options", "", "One or more driver-specific key=value pairs that may be necessary to manage the share server (Optional, Default=None).")
	share_server_adoptCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_server_adoptCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_server_adoptCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_server_adoptCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_server_adoptCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_server_adoptCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_server_adoptCmd.Flags().String("share-network-subnet", "", "Share network subnet where share server has network  allocations in.The default subnet will be used if it's not specified.")
	share_server_adoptCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_server_adoptCmd.Flags().Bool("wait", false, "Wait until share server is adopted")
	share_serverCmd.AddCommand(share_server_adoptCmd)
}
