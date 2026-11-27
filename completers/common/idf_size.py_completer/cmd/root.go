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
	rootCmd.Flags().Bool("dep-rev", false, "Use reverse dependencies for --archive-dependencies")
	rootCmd.Flags().Bool("dep-reverse", false, "Use reverse dependencies for --archive-dependencies")
	rootCmd.Flags().Bool("dep-sym", false, "Include dependency symbols for --archive-dependencies")
	rootCmd.Flags().Bool("dep-symbols", false, "Include dependency symbols for --archive-dependencies")
	rootCmd.Flags().String("diff", "", "Compare sizes with another map file")
	rootCmd.Flags().Bool("files", false, "Print per-file sizes")
	rootCmd.Flags().String("format", "", "Specify output format")
	rootCmd.Flags().StringP("output-file", "o", "", "Print output to the specified file")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"diff":        carapace.ActionFiles(".map"),
		"format":      carapace.ActionValues("csv", "dot", "json2", "raw", "table", "text", "tree"),
		"output-file": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(".map"),
	)
}
