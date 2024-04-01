package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xclip"
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

	rootCmd.Flags().StringN("display", "d", "", "X display to use")
	rootCmd.Flags().BoolN("filter", "f", false, "Print the text piped to standard in back to standard out")
	rootCmd.Flags().BoolN("help", "h", false, "Show quick summary of options")
	rootCmd.Flags().BoolN("in", "i", true, "Read text into X selection")
	rootCmd.Flags().IntN("loops", "l", 0, "Number of X selection requests to wait before exiting")
	rootCmd.Flags().BoolS("noutf8", "noutf8", false, "Operate in legacy mode")
	rootCmd.Flags().BoolN("out", "o", true, "Print the X selection to standard out")
	rootCmd.Flags().BoolS("quiet", "quiet", false, "Run in the foreground, output information messages")
	rootCmd.Flags().BoolN("rmlastnl", "r", false, "Remove any trailing newlines from the selection")
	rootCmd.Flags().StringS("selection", "selection", "primary", "Specify which X selection to use")
	rootCmd.Flags().BoolS("silent", "silent", true, "Fork into background to wait, output errors only")
	rootCmd.Flags().StringN("target", "t", "TEXT", "Specify a specific data format with given target atom")
	rootCmd.Flags().BoolS("verbose", "verbose", false, "Verbose output")
	rootCmd.Flags().BoolS("version", "version", false, "Show version information")

	rootCmd.MarkFlagsMutuallyExclusive("in", "out")
	rootCmd.MarkFlagsMutuallyExclusive("filter", "out")
	rootCmd.MarkFlagsMutuallyExclusive("quiet", "silent", "verbose")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"display":   os.ActionDisplays(),
		"selection": carapace.ActionValues("primary", "secondary", "clipboard"),
		"target": carapace.Batch(
			xclip.ActionTargets(),
			carapace.ActionValues("TARGETS"),
		).ToA(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
