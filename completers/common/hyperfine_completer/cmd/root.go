package cmd

import (
	"slices"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hyperfine",
	Short: "A command-line benchmarking tool.",
	Long:  "https://github.com/sharkdp/hyperfine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("N", "N", false, "An alias for '--shell=none'.")
	rootCmd.Flags().BoolS("S", "S", false, "An alias for '--shell=default'.")
	rootCmd.Flags().StringP("cleanup", "c", "", "Execute CMD after the completion of all benchmarking runs for each individual command to be benchmarked. This is useful if the commands to be benchmarked produce artifacts that need to be cleaned up.")
	rootCmd.Flags().StringSliceP("command-name", "n", nil, "Give a meaningful name to a command. This can be specified multiple times if several commands are benchmarked.")
	rootCmd.Flags().StringSliceP("conclude", "C", nil, "Execute CMD after each timing run. This is useful for killing long-running processes started (e.g. a web server started in --prepare), for example.")
	rootCmd.Flags().StringArray("env", nil, "Set environment variable assignments for benchmarks and their lifecycle commands.")
	rootCmd.Flags().String("export-asciidoc", "", "Export the primary metric summary as an AsciiDoc table.")
	rootCmd.Flags().String("export-csv", "", "Export summary statistics for all selected metrics as CSV.")
	rootCmd.Flags().String("export-json", "", "Export all collected metrics and individual runs as JSON in base units.")
	rootCmd.Flags().String("export-markdown", "", "Export the primary metric summary as a Markdown table.")
	rootCmd.Flags().String("export-orgmode", "", "Export the primary metric summary as an Emacs org-mode table.")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().StringP("ignore-failure", "i", "", "Ignore failures of the benchmarked programs. Without a value or with 'all-non-zero', all non-zero exit codes are ignored. You can also provide a comma-separated list of exit codes to ignore (e.g., --ignore-failure=1,2).")
	rootCmd.Flags().String("input", "", "Control where the input of the benchmark comes from.")
	rootCmd.Flags().StringP("max-runs", "M", "", "Perform at most NUM runs for each command. By default, there is no limit.")
	rootCmd.Flags().String("metrics", "", "Select a metric preset or a comma-separated list of metrics with optional units.")
	rootCmd.Flags().StringP("min-runs", "m", "", "Perform at least NUM runs for each command (default: 10).")
	rootCmd.Flags().StringSlice("output", nil, "Control where the output of the benchmark is redirected. Note that some programs like 'grep' detect when standard output is /dev/null and apply certain optimizations. To avoid that, consider using '--output=pipe'.")
	rootCmd.Flags().StringSliceP("parameter-list", "L", nil, "Perform benchmark runs for each value in the comma-separated list VALUES. Replaces the string '{VAR}' in each command by the current parameter value.")
	rootCmd.Flags().StringP("parameter-scan", "P", "", "Perform benchmark runs for each value in the range MIN..MAX. Replaces the string '{VAR}' in each command by the current parameter value.")
	rootCmd.Flags().StringP("parameter-step-size", "D", "", "This argument requires --parameter-scan to be specified as well. Traverse the range MIN..MAX in steps of DELTA.")
	rootCmd.Flags().StringSliceP("prepare", "p", nil, "Execute CMD before each timing run. This is useful for clearing disk caches, for example.")
	rootCmd.Flags().StringP("runs", "r", "", "Perform exactly NUM runs for each command. If this option is not specified, hyperfine automatically determines the number of runs.")
	rootCmd.Flags().StringP("setup", "s", "", "Execute CMD before each set of timing runs. This is useful for compiling your software with the provided parameters, or to do any other work that should happen once before a series of benchmark runs, not every time as would happen with the --prepare option.")
	rootCmd.Flags().String("shell", "", "Choose the shell for benchmark commands, or use 'none' for direct execution (the default).")
	rootCmd.Flags().Bool("show-output", false, "Print the stdout and stderr of the benchmark instead of suppressing it. This will increase the time it takes for benchmarks to run, so it should only be used for debugging purposes or when trying to benchmark output speed.")
	rootCmd.Flags().String("style", "", "Set output style type (default: auto). Set this to 'basic' to disable output coloring and interactive elements. Set it to 'full' to enable all effects even if no interactive terminal was detected. Set this to 'nocolor' to keep the interactive output without any colors. Set this to 'color' to keep the colors without any interactive output. Set this to 'none' to disable all the output of the tool.")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flags().StringP("warmup", "w", "", "Perform NUM warmup runs before the actual benchmark. This can be used to fill (disk) caches for I/O-heavy programs.")
	rootCmd.Flag("ignore-failure").NoOptDefVal = " "

	rootCmd.Flag("parameter-list").Nargs = -1
	rootCmd.Flag("parameter-scan").Nargs = -1

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cleanup":         bridge.ActionCarapaceBin().Split(),
		"conclude":        bridge.ActionCarapaceBin().Split(),
		"export-asciidoc": carapace.ActionFiles(),
		"export-csv":      carapace.ActionFiles(),
		"export-json":     carapace.ActionFiles(),
		"export-markdown": carapace.ActionFiles(),
		"export-orgmode":  carapace.ActionFiles(),
		"input": carapace.Batch(
			carapace.ActionValuesDescribed("null", "Read from /dev/null"),
			carapace.ActionFiles(),
		).ToA(),
		"metrics": actionMetrics(),
		"output": carapace.Batch(
			carapace.ActionValuesDescribed(
				"inherit", "Don't redirect the output at all",
				"null", "Redirect output to /dev/null",
				"pipe", "Feed the output through a pipe before discarding it",
			),
			carapace.ActionFiles(),
		).ToA(),
		"prepare": bridge.ActionCarapaceBin().Split(),
		"setup":   bridge.ActionCarapaceBin().Split(),
		"shell": carapace.Batch(
			os.ActionShells(),
			carapace.ActionValues("default", "none"),
		).ToA(),
		"style": carapace.ActionValuesDescribed(
			"basic", "disable output coloring and interactive elements",
			"color", "keep the colors without any interactive output",
			"full", "enable all effects",
			"nocolor", "keep the interactive output without any colors",
			"none", "disable all the output of the tool",
		),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		bridge.ActionCarapaceBin().Split(),
	)
}

func actionMetrics() carapace.Action {
	return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
		selected := make([]string, 0, len(c.Parts))
		for _, part := range c.Parts {
			id, _, _ := strings.Cut(part, ":")
			switch id {
			case "default", "time", "all":
				return carapace.ActionValues()
			}
			selected = append(selected, id)
		}
		standalone := len(c.Parts) == 0

		return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 {
				metrics := carapace.ActionValuesDescribed(
					"branch_misses", "Branch misses",
					"cache_misses", "Cache misses",
					"cache_references", "Cache references",
					"cpu_cycles", "CPU cycles",
					"instructions", "Instructions executed",
					"memory_peak_resident", "Peak resident memory",
					"time_cpu", "CPU time",
					"time_system", "System CPU time",
					"time_user", "User CPU time",
					"time_wall_clock", "Wall-clock time",
				)
				if !standalone {
					return metrics.Filter(selected...)
				}
				return carapace.Batch(metrics, carapace.ActionValuesDescribed(
					"all", "Preset: all available metrics",
					"default", "Preset: wall-clock time and peak memory",
					"time", "Preset: wall-clock, CPU, user, and system time",
				)).ToA()
			}

			if slices.Contains(selected, c.Parts[0]) {
				return carapace.ActionValues()
			}
			switch c.Parts[0] {
			case "time_wall_clock", "time_cpu", "time_user", "time_system":
				return carapace.ActionValues( // carapace:curated
					"ns", "us", "ms", "s", "min", "h",
				)
			case "memory_peak_resident":
				return carapace.ActionValues( // carapace:curated
					"B", "kB", "MB", "GB", "TB",
					"KiB", "MiB", "GiB", "TiB",
				)
			case "cpu_cycles", "instructions", "cache_references", "cache_misses", "branch_misses":
				return carapace.ActionValues( // carapace:curated
					"count", "k", "M", "B",
				)
			default:
				return carapace.ActionValues()
			}
		})
	})
}
