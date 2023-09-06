package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var top_nodeCmd = &cobra.Command{
	Use:     "node [NAME | -l label]",
	Short:   "Display resource (CPU/memory) usage of nodes",
	Aliases: []string{"nodes", "no"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(top_nodeCmd).Standalone()

	top_nodeCmd.Flags().Bool("no-headers", false, "If present, print output without headers")
	top_nodeCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
	top_nodeCmd.Flags().Bool("show-capacity", false, "Print node resources based on Capacity instead of Allocatable(default) of the nodes.")
	top_nodeCmd.Flags().String("sort-by", "", "If non-empty, sort nodes list using specified field. The field can be either 'cpu' or 'memory'.")
	top_nodeCmd.Flags().Bool("use-protocol-buffers", false, "Enables using protocol-buffers to access Metrics API.")
	topCmd.AddCommand(top_nodeCmd)

	carapace.Gen(top_nodeCmd).FlagCompletion(carapace.ActionMap{
		"sort-by": carapace.ActionValues("cpu", "memory"),
	})

	carapace.Gen(top_nodeCmd).PositionalCompletion(
		kubectl.ActionResources(kubectl.ResourceOpts{Namespace: "", Types: "nodes"}),
	)
}
