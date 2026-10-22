package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_host_failoverCmd = &cobra.Command{
	Use:   "failover",
	Short: "Failover volume host to different backend",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_host_failoverCmd).Standalone()

	volume_host_failoverCmd.Flags().String("volume-backend", "", "The ID of the volume backend replication target where the host will failover to (required)")
	volume_host_failoverCmd.MarkFlagRequired("volume-backend")
	volume_hostCmd.AddCommand(volume_host_failoverCmd)
}
