package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_unshelveCmd = &cobra.Command{
	Use:   "unshelve",
	Short: "Unshelve server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_unshelveCmd).Standalone()

	server_unshelveCmd.Flags().String("availability-zone", "", "Name of the availability zone in which to unshelve a SHELVED_OFFLOADED server (supported by --os-compute-api-version 2.77 or above)")
	server_unshelveCmd.Flags().String("host", "", "Name of the destination host in which to unshelve a SHELVED_OFFLOADED server (supported by --os-compute-api-version 2.91 or above)")
	server_unshelveCmd.Flags().Bool("no-availability-zone", false, "Unpin the availability zone of a SHELVED_OFFLOADED server.")
	server_unshelveCmd.Flags().Bool("wait", false, "Wait for unshelve operation to complete")
	serverCmd.AddCommand(server_unshelveCmd)
}
