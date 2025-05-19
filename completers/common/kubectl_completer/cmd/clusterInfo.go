package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clusterInfoCmd = &cobra.Command{
	Use:     "cluster-info",
	Short:   "Display cluster information",
	GroupID: "cluster management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clusterInfoCmd).Standalone()

	rootCmd.AddCommand(clusterInfoCmd)
}
