package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/k3d"
	"github.com/spf13/cobra"
)

var cluster_stopCmd = &cobra.Command{
	Use:   "stop [NAME [NAME...] | --all]",
	Short: "Stop existing k3d cluster(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_stopCmd).Standalone()

	cluster_stopCmd.Flags().BoolP("all", "a", false, "Stop all existing clusters")
	clusterCmd.AddCommand(cluster_stopCmd)

	carapace.Gen(cluster_stopCmd).PositionalAnyCompletion(
		k3d.ActionClusters().FilterArgs(),
	)
}
