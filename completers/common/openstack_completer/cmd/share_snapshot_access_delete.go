package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_snapshot_access_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete access to a snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_snapshot_access_deleteCmd).Standalone()

	share_snapshot_accessCmd.AddCommand(share_snapshot_access_deleteCmd)
}
