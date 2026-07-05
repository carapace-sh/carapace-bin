package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qpreferrednodeCmd = &cobra.Command{
	Use:   "qpreferrednode",
	Short: "query the preferred NUMA node of a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qpreferrednodeCmd).Standalone()
	rootCmd.AddCommand(qpreferrednodeCmd)
}
