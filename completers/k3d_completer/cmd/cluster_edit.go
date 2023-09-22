package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cluster_editCmd = &cobra.Command{
	Use:     "edit CLUSTER",
	Short:   "[EXPERIMENTAL] Edit cluster(s).",
	Aliases: []string{"update"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_editCmd).Standalone()

	cluster_editCmd.Flags().StringSlice("port-add", []string{}, "[EXPERIMENTAL] Map ports from the node containers (via the serverlb) to the host (Format: `[HOST:][HOSTPORT:]CONTAINERPORT[/PROTOCOL][@NODEFILTER]`)")
	clusterCmd.AddCommand(cluster_editCmd)
}
