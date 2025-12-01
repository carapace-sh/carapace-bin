package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
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
	top_nodeCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	top_nodeCmd.Flags().Bool("show-capacity", false, "Print node resources based on Capacity instead of Allocatable(default) of the nodes.")
	top_nodeCmd.Flags().Bool("show-swap", false, "Print node resources related to swap memory.")
	top_nodeCmd.Flags().String("sort-by", "", "If non-empty, sort nodes list using specified field. The field can be either 'cpu' or 'memory'.")
	top_nodeCmd.Flags().Bool("use-protocol-buffers", false, "Enables using protocol-buffers to access Metrics API.")
	topCmd.AddCommand(top_nodeCmd)

	carapace.Gen(top_nodeCmd).FlagCompletion(carapace.ActionMap{
		"sort-by": carapace.ActionValues("cpu", "memory"),
	})

	carapace.Gen(top_nodeCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionResources(kubectl.ResourceOpts{
				Context:   rootCmd.Flag("context").Value.String(),
				Namespace: rootCmd.Flag("namespace").Value.String(),
				Types:     "nodes",
			})
		}),
	)
}
