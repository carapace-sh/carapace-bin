package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/k3d"
	"github.com/spf13/cobra"
)

var node_stopCmd = &cobra.Command{
	Use:   "stop NAME",
	Short: "Stop an existing k3d node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_stopCmd).Standalone()

	nodeCmd.AddCommand(node_stopCmd)

	carapace.Gen(node_stopCmd).PositionalCompletion(
		k3d.ActionNodes(k3d.NodeOpts{Running: true}),
	)
}
