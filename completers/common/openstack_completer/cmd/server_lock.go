package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Lock server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_lockCmd).Standalone()

	server_lockCmd.Flags().String("reason", "", "Reason for locking the server(s) (supported by --os-compute-api-version 2.73 or above)")
	serverCmd.AddCommand(server_lockCmd)
}
