package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_nodePool_replaceNodeCmd = &cobra.Command{
	Use:   "replace-node",
	Short: "Replace node with a new one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_nodePool_replaceNodeCmd).Standalone()
	kubernetes_cluster_nodePool_replaceNodeCmd.Flags().BoolP("force", "f", false, "Replace node without confirmation prompt")
	kubernetes_cluster_nodePool_replaceNodeCmd.Flags().Bool("skip-drain", false, "Skip draining the node before replacement")
	kubernetes_cluster_nodePoolCmd.AddCommand(kubernetes_cluster_nodePool_replaceNodeCmd)
}
