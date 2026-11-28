package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "idf_size.py",
	Short: "Display firmware size information for ESP-IDF projects",
	Long:  "https://github.com/espressif/esp-idf-size",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("archive-dependencies", false, "Display dependencies for all archives")
	rootCmd.Flags().Bool("archive-deps", false, "Display dependencies for all archives")
	rootCmd.Flags().String("archive-details", "", "Print detailed symbols per archive")
	rootCmd.Flags().Bool("archives", false, "Print per-archive sizes")
	rootCmd.Flags().BoolP("debug", "d", false, "Print debug information")
	rootCmd.Flags().Bool("dep-rev", false, "Use reverse dependencies for --archive-dependencies")
	rootCmd.Flags().Bool("dep-reverse", false, "Use reverse dependencies for --archive-dependencies")
	rootCmd.Flags().Bool("dep-sym", false, "Include dependency symbols for --archive-dependencies")
	rootCmd.Flags().Bool("dep-symbols", false, "Include dependency symbols for --archive-dependencies")
	rootCmd.Flags().String("diff", "", "Compare sizes with another map file")
	rootCmd.Flags().Bool("doc", false, "Display more comprehensive documentation")
	rootCmd.Flags().Bool("files", false, "Print per-file sizes")
	rootCmd.Flags().StringArrayP("filter", "F", nil, "Filter archives, object files, or symbols by pattern")
	rootCmd.Flags().Bool("force-terminal", false, "Enable terminal control codes even if output is not a terminal")
	rootCmd.Flags().String("format", "", "Specify output format")
	rootCmd.Flags().Bool("lto", false, "Enable usage of DWARF information to identify archives for LTO object files")
	rootCmd.Flags().Bool("no-abbrev", false, "Do not abbreviate section and file names")
	rootCmd.Flags().Bool("no-color", false, "Disable ANSI color escape sequences")
	rootCmd.Flags().Bool("no-lto", false, "Disable usage of DWARF information to identify archives for LTO object files")
	rootCmd.Flags().StringP("output-file", "o", "", "Print output to the specified file")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress all output")
	rootCmd.Flags().Bool("show-unchanged", false, "Show unchanged items for --diff operation")
	rootCmd.Flags().Bool("show-unused", false, "Show unused memory types and sections")
	rootCmd.Flags().StringP("sort", "s", "", "Sort table rows based on specified column")
	rootCmd.Flags().Bool("sort-diff", false, "Sort entries based on diff value instead of size")
	rootCmd.Flags().Bool("sort-reverse", false, "Sort entries in reversed order")
	rootCmd.Flags().Bool("unify", false, "Use abbreviated names with aggregated size information")
	rootCmd.Flags().Bool("use-flash-size", false, "Show the total flash size as defined in the link map file")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"diff":        carapace.ActionFiles(".map"),
		"format":      carapace.ActionValues("csv", "dot", "json2", "raw", "table", "text", "tree"),
		"output-file": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(".map"),
	)
}
