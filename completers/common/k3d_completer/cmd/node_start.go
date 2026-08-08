package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/k3d"
	"github.com/spf13/cobra"
)

var node_startCmd = &cobra.Command{
	Use:   "start NODE",
	Short: "Start an existing k3d node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_startCmd).Standalone()

	nodeCmd.AddCommand(node_startCmd)

	carapace.Gen(node_startCmd).PositionalCompletion(
		k3d.ActionNodes(k3d.NodeOpts{Stopped: true}),
	)
}
