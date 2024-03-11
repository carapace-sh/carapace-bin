package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var node_rmCmd = &cobra.Command{
	Use:     "rm [OPTIONS] NODE [NODE...]",
	Short:   "Remove one or more nodes from the swarm",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_rmCmd).Standalone()

	node_rmCmd.Flags().BoolP("force", "f", false, "Force remove a node from the swarm")
	nodeCmd.AddCommand(node_rmCmd)

	carapace.Gen(node_rmCmd).PositionalAnyCompletion(docker.ActionNodes())
}
