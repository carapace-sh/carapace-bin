package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:     "service",
	Short:   "Manage Swarm services",
	GroupID: "swarm",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceCmd).Standalone()

	rootCmd.AddCommand(serviceCmd)
}
