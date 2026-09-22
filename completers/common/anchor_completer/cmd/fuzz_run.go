package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fuzz_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a fuzz test",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fuzz_runCmd).Standalone()

	fuzz_runCmd.Flags().String("binary-in", "", "Run a prebuilt harness binary directly")
	fuzz_runCmd.Flags().StringP("cores", "j", "", "Run N parallel fuzzer workers")
	fuzz_runCmd.Flags().String("corpus-in", "", "Load seed corpus from directory")
	fuzz_runCmd.Flags().String("corpus-out", "", "Write corpus to directory")
	fuzz_runCmd.Flags().Bool("coverage", false, "Enable coverage reporting (single-core only)")
	fuzz_runCmd.Flags().String("crashes-meta-out", "", "Custom directory for .meta.json files (default: same as crashes_out)")
	fuzz_runCmd.Flags().String("crashes-out", "", "Custom crash output directory")
	fuzz_runCmd.Flags().Bool("dry-run", false, "Validate setup without fuzzing")
	fuzz_runCmd.Flags().BoolP("help", "h", false, "Print help")
	fuzz_runCmd.Flags().String("lcov-out", "", "LCOV coverage output path")
	fuzz_runCmd.Flags().String("max-actions", "", "Maximum number of actions per fuzzer iteration (default: 8 stateless, 100 stateful)")
	fuzz_runCmd.Flags().String("max-depth", "", "Maximum state depth (action chain length) in stateful mode (default: 15)")
	fuzz_runCmd.Flags().String("mode", "", "Remote fuzzing operational mode (dry_run, explore, coverage, reproduce, corpus_merge)")
	fuzz_runCmd.Flags().Bool("no-tracing", false, "Disable SVM register tracing for higher throughput (no coverage guidance)")
	fuzz_runCmd.Flags().String("pool-size", "", "State pool capacity in stateful mode (default: 256000)")
	fuzz_runCmd.Flags().String("program-so", "", "Path to alternative program .so binary")
	fuzz_runCmd.Flags().Bool("release", false, "Build in release mode")
	fuzz_runCmd.Flags().String("replay", "", "Replay a single crash/input file")
	fuzz_runCmd.Flags().String("seed", "", "Random seed for reproducible fuzzing")
	fuzz_runCmd.Flags().Bool("stateful", false, "Stateful fuzzing: single action per iteration with state pool")
	fuzz_runCmd.Flags().Bool("stats", false, "Show detailed performance stats (profiling, memory, pick/exec breakdown)")
	fuzz_runCmd.Flags().Bool("stop-on-crash", false, "Stop fuzzing on first crash")
	fuzz_runCmd.Flags().String("symbols", "", "Path to debug binary with DWARF symbols (for source-level coverage with --coverage)")
	fuzz_runCmd.Flags().String("timeout", "", "Stop after N seconds")
	fuzzCmd.AddCommand(fuzz_runCmd)

	carapace.Gen(fuzz_runCmd).FlagCompletion(carapace.ActionMap{
		"binary-in":        carapace.ActionFiles(),
		"corpus-in":        carapace.ActionFiles(),
		"corpus-out":       carapace.ActionFiles(),
		"crashes-meta-out": carapace.ActionFiles(),
		"crashes-out":      carapace.ActionFiles(),
		"lcov-out":         carapace.ActionFiles(),
		"program-so":       carapace.ActionFiles(".so"),
		"replay":           carapace.ActionFiles(),
		"symbols":          carapace.ActionFiles(),
	})
}
