package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_server_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete one or more share servers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_server_deleteCmd).Standalone()

	share_server_deleteCmd.Flags().Bool("wait", false, "Wait for share server deletion.")
	share_serverCmd.AddCommand(share_server_deleteCmd)
}
