package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_group_failoverCmd = &cobra.Command{
	Use:   "failover",
	Short: "Failover replication for a volume group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_failoverCmd).Standalone()

	volume_group_failoverCmd.Flags().Bool("allow-attached-volume", false, "Allow group with attached volumes to be failed over.")
	volume_group_failoverCmd.Flags().Bool("disallow-attached-volume", false, "Disallow group with attached volumes to be failed over.")
	volume_group_failoverCmd.Flags().String("secondary-backend-id", "", "Secondary backend ID.")
	volume_groupCmd.AddCommand(volume_group_failoverCmd)
}
