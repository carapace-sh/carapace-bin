package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fsck_hfs",
	Short: "HFS+ filesystem consistency check",
	Long:  "https://keith.github.io/xcode-manpages/fsck_hfs.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("E", "E", false, "E flag")
	rootCmd.Flags().BoolS("S", "S", false, "Snapshot")
	rootCmd.Flags().BoolS("d", "d", false, "Debug output")
	rootCmd.Flags().BoolS("f", "f", false, "Force")
	rootCmd.Flags().BoolS("g", "g", false, "G flag")
	rootCmd.Flags().BoolS("l", "l", false, "Live filesystem")
	rootCmd.Flags().BoolS("n", "n", false, "Assume no to all questions")
	rootCmd.Flags().BoolS("p", "p", false, "Preen mode (automatic repair)")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode")
	rootCmd.Flags().BoolS("r", "r", false, "Rebuild catalog B-tree")
	rootCmd.Flags().BoolS("x", "x", false, "Extended mode")
	rootCmd.Flags().BoolS("y", "y", false, "Assume yes to all questions")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
