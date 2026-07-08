package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "benchmark",
	Short: "Benchmark execution of a callable",
	Long:  "https://elv.sh/ref/builtin.html#benchmark",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("min-runs", "min-runs", "5", "minimum number of runs")
	rootCmd.Flags().StringS("min-time", "min-time", "1s", "minimum duration")
	rootCmd.Flags().StringS("on-end", "on-end", "", "callback with stats on end")
	rootCmd.Flags().StringS("on-run-end", "on-run-end", "", "callback after each run")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("callable to benchmark"),
	)
}
