package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/k3d"
	"github.com/spf13/cobra"
)

var node_listCmd = &cobra.Command{
	Use:     "list [NODE [NODE...]]",
	Short:   "List node(s)",
	Aliases: []string{"ls", "get"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_listCmd).Standalone()

	node_listCmd.Flags().Bool("no-headers", false, "Disable headers")
	node_listCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml")
	nodeCmd.AddCommand(node_listCmd)

	carapace.Gen(node_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("json", "yaml"),
	})

	carapace.Gen(node_listCmd).PositionalAnyCompletion(
		k3d.ActionNodes(k3d.NodeOpts{}.Default()).FilterArgs(),
	)
}
