package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var benchmarkCmd = &cobra.Command{
	Use:   "benchmark",
	Short: "run a performance benchmark",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(benchmarkCmd).Standalone()
	rootCmd.AddCommand(benchmarkCmd)
}
