package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_shelveCmd = &cobra.Command{
	Use:   "shelve",
	Short: "Shelve and optionally offload server(s).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_shelveCmd).Standalone()

	server_shelveCmd.Flags().Bool("offload", false, "Remove the shelved server(s) from the host (admin only).")
	server_shelveCmd.Flags().Bool("wait", false, "Wait for shelve and/or offload operation to complete")
	serverCmd.AddCommand(server_shelveCmd)
}
