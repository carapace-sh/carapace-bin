package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var node_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops a node in a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_stopCmd).Standalone()
	nodeCmd.AddCommand(node_stopCmd)
}
