package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var analysisStatsCmd = &cobra.Command{
	Use:   "analysis-stats",
	Short: "Batch typecheck project and print summary statistics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(analysisStatsCmd).Standalone()

	analysisStatsCmd.Flags().Bool("disable-build-scripts", false, "Don't run build scripts or load `OUT_DIR` values by running `cargo check` before analysis")
	analysisStatsCmd.Flags().Bool("disable-proc-macros", false, "Don't use expand proc macros")
	analysisStatsCmd.Flags().Bool("memory-usage", false, "Collect memory usage statistics")
	analysisStatsCmd.Flags().Bool("no-sysroot", false, "Don't load sysroot crates (`std`, `core` & friends)")
	analysisStatsCmd.Flags().StringP("only", "o", "", "Only analyze items matching this path")
	analysisStatsCmd.Flags().String("output", "", "")
	analysisStatsCmd.Flags().Bool("parallel", false, "Run type inference in parallel")
	analysisStatsCmd.Flags().Bool("randomize", false, "Randomize order in which crates, modules, and items are processed")
	analysisStatsCmd.Flags().Bool("skip-inference", false, "Only resolve names, don't run type inference")
	analysisStatsCmd.Flags().Bool("source-stats", false, "Print the total length of all source and macro files (whitespace is not counted)")
	analysisStatsCmd.Flags().Bool("with-deps", false, "Also analyze all dependencies")
	rootCmd.AddCommand(analysisStatsCmd)

	carapace.Gen(analysisStatsCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
