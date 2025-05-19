package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:     "node",
	Short:   "Manage Swarm nodes",
	GroupID: "swarm",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nodeCmd).Standalone()

	rootCmd.AddCommand(nodeCmd)
}
