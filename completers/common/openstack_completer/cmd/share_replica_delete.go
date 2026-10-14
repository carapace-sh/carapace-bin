package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_replica_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete one or more share replicas",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_replica_deleteCmd).Standalone()

	share_replica_deleteCmd.Flags().Bool("force", false, "Attempt to force delete a replica on its backend.")
	share_replica_deleteCmd.Flags().Bool("wait", false, "Wait for share replica deletion")
	share_replicaCmd.AddCommand(share_replica_deleteCmd)
}
