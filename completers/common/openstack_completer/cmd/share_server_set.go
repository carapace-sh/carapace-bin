package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_server_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set share server properties (Admin only).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_server_setCmd).Standalone()

	share_server_setCmd.Flags().String("status", "", "Assign a status to the share server.")
	share_server_setCmd.Flags().String("task-state", "", "Indicate which task state to assign the share server.")
	share_serverCmd.AddCommand(share_server_setCmd)
}
