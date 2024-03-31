package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/xclip_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xclip [flags] [FILE]",
	Short: "command line interface to X selections",
	Long:  "https://github.com/astrand/xclip/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("display", "d", "", "X display to use")
	rootCmd.Flags().StringS("display-long", "display", "", "X display to use")
	rootCmd.Flags().BoolS("filter", "f", false, "Print the text piped to standard in back to standard out")
	rootCmd.Flags().BoolS("filter-long", "filter", false, "Print the text piped to standard in back to standard out")
	rootCmd.Flags().BoolP("help", "h", false, "Show quick summary of options")
	rootCmd.Flags().BoolS("input", "i", true, "Read text into X selection")
	rootCmd.Flags().BoolS("input-long", "in", true, "Read text into X selection")
	rootCmd.Flags().IntS("loops", "l", 0, "Number of X selection requests to wait before exiting")
	rootCmd.Flags().IntS("loops-long", "loops", 0, "Number of X selection requests to wait before exiting")
	rootCmd.Flags().BoolS("noutf8", "noutf8", false, "Operate in legacy mode")
	rootCmd.Flags().BoolS("output", "o", true, "Print the X selection to standard out")
	rootCmd.Flags().BoolS("output-long", "out", true, "Print the X selection to standard out")
	rootCmd.Flags().BoolS("quiet", "quiet", false, "Run in the foreground, output information messages")
	rootCmd.Flags().BoolS("rmlast", "r", false, "Remove any trailing newlines from the selection")
	rootCmd.Flags().BoolS("rmlast-long", "rmlastnl", false, "Remove any trailing newlines from the selection")
	rootCmd.Flags().StringS("selection", "selection", "primary", "Specify which X selection to use")
	rootCmd.Flags().BoolS("silent", "silent", true, "Fork into background to wait, output errors only")
	rootCmd.Flags().StringS("target", "t", "TEXT", "Specify a specific data format with given target atom")
	rootCmd.Flags().StringS("target-long", "target", "TEXT", "Specify a specific data format with given target atom")
	rootCmd.Flags().BoolS("verbose", "verbose", false, "Verbose output")
	rootCmd.Flags().BoolP("version", "version", false, "Show version information")

	rootCmd.MarkFlagsMutuallyExclusive("display", "display-long")
	rootCmd.MarkFlagsMutuallyExclusive("input", "input-long", "output", "output-long")
	rootCmd.MarkFlagsMutuallyExclusive("filter", "filter-long", "output", "output-long")
	rootCmd.MarkFlagsMutuallyExclusive("loops", "loops-long")
	rootCmd.MarkFlagsMutuallyExclusive("quiet", "silent", "verbose")
	rootCmd.MarkFlagsMutuallyExclusive("rmlast", "rmlast-long")
	rootCmd.MarkFlagsMutuallyExclusive("target", "target-long")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"selection": carapace.ActionValues("primary", "secondary", "clipboard"),
		"target":    xclip.ActionTarget(),
	})
	carapace.Gen(rootCmd).PositionalCompletion(carapace.ActionFiles())
}
