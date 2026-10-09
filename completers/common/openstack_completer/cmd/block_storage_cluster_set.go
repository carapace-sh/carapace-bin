package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var block_storage_cluster_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set block storage cluster properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(block_storage_cluster_setCmd).Standalone()

	block_storage_cluster_setCmd.Flags().String("binary", "", "Name of binary to filter by; defaults to 'cinder-volume' (optional)")
	block_storage_cluster_setCmd.Flags().Bool("disable", false, "Disable cluster")
	block_storage_cluster_setCmd.Flags().String("disable-reason", "", "Reason for disabling the cluster (should be used with --disable option)")
	block_storage_cluster_setCmd.Flags().Bool("enable", false, "Enable cluster")
	block_storage_clusterCmd.AddCommand(block_storage_cluster_setCmd)
}
