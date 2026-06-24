package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var refs_optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize ref storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refs_optimizeCmd).Standalone()

	refsCmd.AddCommand(refs_optimizeCmd)
}
