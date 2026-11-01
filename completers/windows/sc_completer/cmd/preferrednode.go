package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var preferrednodeCmd = &cobra.Command{
	Use:   "preferrednode",
	Short: "set the preferred NUMA node for a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(preferrednodeCmd).Standalone()
	rootCmd.AddCommand(preferrednodeCmd)
}
