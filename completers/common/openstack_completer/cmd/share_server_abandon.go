package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_server_abandonCmd = &cobra.Command{
	Use:   "abandon",
	Short: "Remove one or more share server(s) (Admin only).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_server_abandonCmd).Standalone()

	share_server_abandonCmd.Flags().Bool("force", false, "Enforces the unmanage share server operation, even if the backend driver does not support it.")
	share_server_abandonCmd.Flags().Bool("wait", false, "Wait until share server is abandoned")
	share_serverCmd.AddCommand(share_server_abandonCmd)
}
