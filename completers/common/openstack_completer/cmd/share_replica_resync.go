package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_replica_resyncCmd = &cobra.Command{
	Use:   "resync",
	Short: "Attempt to update the share replica with its 'active' mirror.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_replica_resyncCmd).Standalone()

	share_replicaCmd.AddCommand(share_replica_resyncCmd)
}
