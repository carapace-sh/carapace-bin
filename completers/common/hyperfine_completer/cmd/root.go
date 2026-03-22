package cmd

import (
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
	rootCmd.Flags().StringP("cleanup", "c", "", "Execute CMD after the completion of all benchmarking runs for each individual command to be benchmarked. This is useful if the commands to be benchmarked produce artifacts that need to be cleaned up.")
	rootCmd.Flags().StringSliceP("command-name", "n", nil, "Give a meaningful name to a command. This can be specified multiple times if several commands are benchmarked.")
	rootCmd.Flags().StringSliceP("conclude", "C", nil, "Execute CMD after each timing run. This is useful for killing long-running processes started (e.g. a web server started in --prepare), for example.")
	rootCmd.Flags().String("export-asciidoc", "", "Export the timing summary statistics as an AsciiDoc table to the given FILE. The output time unit can be changed using the --time-unit option.")
	rootCmd.Flags().String("export-csv", "", "Export the timing summary statistics as CSV to the given FILE. If you need the timing results for each individual run, use the JSON export format. The output time unit is always seconds.")
	rootCmd.Flags().String("export-json", "", "Export the timing summary statistics and timings of individual runs as JSON to the given FILE. The output time unit is always seconds")
	rootCmd.Flags().String("export-markdown", "", "Export the timing summary statistics as a Markdown table to the given FILE. The output time unit can be changed using the --time-unit option.")
	rootCmd.Flags().String("export-orgmode", "", "Export the timing summary statistics as an Emacs org-mode table to the given FILE. The output time unit can be changed using the --time-unit option.")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().StringP("ignore-failure", "i", "", "Ignore failures of the benchmarked programs. Without a value or with 'all-non-zero', all non-zero exit codes are ignored. You can also provide a comma-separated list of exit codes to ignore (e.g., --ignore-failure=1,2).")
	rootCmd.Flags().String("input", "", "Control where the input of the benchmark comes from.")
	rootCmd.Flags().StringP("max-runs", "M", "", "Perform at most NUM runs for each command. By default, there is no limit.")
	rootCmd.Flags().StringP("min-runs", "m", "", "Perform at least NUM runs for each command (default: 10).")
	rootCmd.Flags().StringSlice("output", nil, "Control where the output of the benchmark is redirected. Note that some programs like 'grep' detect when standard output is /dev/null and apply certain optimizations. To avoid that, consider using '--output=pipe'.")
	rootCmd.Flags().StringSliceP("parameter-list", "L", nil, "Perform benchmark runs for each value in the comma-separated list VALUES. Replaces the string '{VAR}' in each command by the current parameter value.")
	rootCmd.Flags().StringP("parameter-scan", "P", "", "Perform benchmark runs for each value in the range MIN..MAX. Replaces the string '{VAR}' in each command by the current parameter value.")
	rootCmd.Flags().StringP("parameter-step-size", "D", "", "This argument requires --parameter-scan to be specified as well. Traverse the range MIN..MAX in steps of DELTA.")
	rootCmd.Flags().StringSliceP("prepare", "p", nil, "Execute CMD before each timing run. This is useful for clearing disk caches, for example.")
	rootCmd.Flags().String("reference", "", "The reference command for the relative comparison of results. If this is unset, results are compared with the fastest command as reference.")
	rootCmd.Flags().String("reference-name", "", "Give a meaningful name to the reference command.")
	rootCmd.Flags().StringP("runs", "r", "", "Perform exactly NUM runs for each command. If this option is not specified, hyperfine automatically determines the number of runs.")
	rootCmd.Flags().StringP("setup", "s", "", "Execute CMD before each set of timing runs. This is useful for compiling your software with the provided parameters, or to do any other work that should happen once before a series of benchmark runs, not every time as would happen with the --prepare option.")
	rootCmd.Flags().StringP("shell", "S", "", "Set the shell to use for executing benchmarked commands. This can be the name or the path to the shell executable, or a full command line like \"bash --norc\". It can also be set to \"default\" to explicitly select the default shell on this platform. Finally, this can also be set to \"none\" to disable the shell. In this case, commands will be executed directly. They can still have arguments, but more complex things like \"sleep 0.1; sleep 0.2\" are not possible without a shell.")
	rootCmd.Flags().Bool("show-output", false, "Print the stdout and stderr of the benchmark instead of suppressing it. This will increase the time it takes for benchmarks to run, so it should only be used for debugging purposes or when trying to benchmark output speed.")
	rootCmd.Flags().String("sort", "", "Specify the sort order of the speed comparison summary and the exported tables for markup formats (Markdown, AsciiDoc, org-mode):")
	rootCmd.Flags().String("style", "", "Set output style type (default: auto). Set this to 'basic' to disable output coloring and interactive elements. Set it to 'full' to enable all effects even if no interactive terminal was detected. Set this to 'nocolor' to keep the interactive output without any colors. Set this to 'color' to keep the colors without any interactive output. Set this to 'none' to disable all the output of the tool.")
	rootCmd.Flags().StringP("time-unit", "u", "", "Set the time unit to be used. Possible values: microsecond, millisecond, second. If the option is not given, the time unit is determined automatically. This option affects the standard output as well as all export formats except for CSV and JSON.")
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
		"output": carapace.Batch(
			carapace.ActionValuesDescribed(
				"null", "Redirect output to /dev/null",
				"pipe", "Feed the output through a pipe before discarding it",
				"inherit", "Don't redirect the output at all",
			),
			carapace.ActionFiles(),
		).ToA(),
		"prepare":   bridge.ActionCarapaceBin().Split(),
		"reference": bridge.ActionCarapaceBin().Split(),
		"setup":     bridge.ActionCarapaceBin().Split(),
		"shell": carapace.Batch(
			os.ActionShells(),
			carapace.ActionValues("none"),
		).ToA(),
		"sort": carapace.ActionValuesDescribed(
			"auto", "order speed by time and markup by command",
			"command", "order benchmarks in the way they were specified",
			"mean-time", "order benchmarks by mean runtime",
		),
		"style": carapace.ActionValuesDescribed(
			"basic", "disable output coloring and interactive elements",
			"full", "enable all effects",
			"color", "keep the colors without any interactive output",
			"none", "disable all the output of the tool",
		),
		"time-unit": carapace.ActionValues("microsecond", "millisecond", "second"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		bridge.ActionCarapaceBin().Split(),
	)
}
