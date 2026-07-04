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
	analysisStatsCmd.Flags().Bool("no-test", false, "Don't set #[cfg(test)]")
	analysisStatsCmd.Flags().StringP("only", "o", "", "Only analyze items matching this path")
	analysisStatsCmd.Flags().String("output", "", "")
	analysisStatsCmd.Flags().Bool("parallel", false, "Run type inference in parallel")
	analysisStatsCmd.Flags().String("proc-macro-srv", "", "Run the proc-macro-srv binary at the specified path")
	analysisStatsCmd.Flags().Bool("randomize", false, "Randomize order in which crates, modules, and items are processed")
	analysisStatsCmd.Flags().Bool("run-all-ide-things", false, "Runs several IDE features after analysis")
	analysisStatsCmd.Flags().Bool("run-term-search", false, "Run term search on all the tail expressions")
	analysisStatsCmd.Flags().Bool("skip-const-eval", false, "Skip const evaluation")
	analysisStatsCmd.Flags().Bool("skip-data-layout", false, "Skip data layout calculation")
	analysisStatsCmd.Flags().Bool("skip-inference", false, "Only resolve names, don't run type inference")
	analysisStatsCmd.Flags().Bool("skip-lang-items", false, "Skip lang items fetching")
	analysisStatsCmd.Flags().Bool("skip-lowering", false, "Skip body lowering")
	analysisStatsCmd.Flags().Bool("skip-mir-stats", false, "Skip lowering to mir")
	analysisStatsCmd.Flags().Bool("source-stats", false, "Print the total length of all source and macro files (whitespace is not counted)")
	analysisStatsCmd.Flags().Bool("validate-term-search", false, "Validate term search by running `cargo check` on every response")
	analysisStatsCmd.Flags().Bool("with-deps", false, "Also analyze all dependencies")
	rootCmd.AddCommand(analysisStatsCmd)

	carapace.Gen(analysisStatsCmd).FlagCompletion(carapace.ActionMap{
		"output":         carapace.ActionFiles(),
		"proc-macro-srv": carapace.ActionFiles(),
	})

	carapace.Gen(analysisStatsCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
