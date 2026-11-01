package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "footprint",
	Short: "gather memory information about one or more processes",
	Long:  "https://keith.github.io/xcode-manpages/footprint.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "Target all processes")
	rootCmd.Flags().StringS("f", "f", "", "Format")
	rootCmd.Flags().BoolS("h", "h", false, "Help")
	rootCmd.Flags().StringS("j", "j", "", "JSON output path")
	rootCmd.Flags().StringS("p", "p", "", "Process name or PID")
	rootCmd.Flags().BoolS("s", "s", false, "Sort")
	rootCmd.Flags().String("sort", "", "Sort column")
	rootCmd.Flags().String("swapped", "", "Swapped memory")
	rootCmd.Flags().BoolS("t", "t", false, "Track")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")
	rootCmd.Flags().BoolS("w", "w", false, "Wired")
	rootCmd.Flags().String("wired", "", "Wired memory")
	rootCmd.Flags().StringS("x", "x", "", "Exclude process name or PID")
	rootCmd.Flags().BoolS("y", "y", false, "Swapped")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionValues("bytes", "formatted", "pages"),
		"j": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		ps.ActionProcessIds(),
	)
}
