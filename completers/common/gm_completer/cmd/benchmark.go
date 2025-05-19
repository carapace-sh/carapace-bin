package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var benchmarkCmd = &cobra.Command{
	Use:   "benchmark",
	Short: "benchmark one of the other command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(benchmarkCmd).Standalone()

	benchmarkCmd.Flags().BoolS("concurrent", "concurrent", false, "run multiple commands in parallel")
	benchmarkCmd.Flags().IntS("duration", "duration", 0, "duration to run benchmark (in seconds)")
	benchmarkCmd.Flags().IntS("iterations", "iterations", 0, "number of command iterations per benchmark")
	benchmarkCmd.Flags().BoolS("rawcsv", "rawcsv", false, "CSV output (threads,iterations,user_time,elapsed_time)")
	benchmarkCmd.Flags().IntS("stepthreads", "stepthreads", 0, "step benchmark with increasing number of threads")
	rootCmd.AddCommand(benchmarkCmd)

	carapace.Gen(benchmarkCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
