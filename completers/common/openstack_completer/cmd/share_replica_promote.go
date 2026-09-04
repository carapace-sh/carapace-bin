package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_replica_promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "Promote specified replica to 'active' replica_state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_replica_promoteCmd).Standalone()

	share_replica_promoteCmd.Flags().String("quiesce-wait-time", "", "Quiesce wait time in seconds.")
	share_replica_promoteCmd.Flags().Bool("wait", false, "Wait for share replica promotion")
	share_replicaCmd.AddCommand(share_replica_promoteCmd)
}
