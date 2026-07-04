package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "find",
	Short: "walk a file hierarchy",
	Long:  "https://keith.github.io/xcode-manpages/find.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("E", "E", false, "Interpret regular expressions using extended regex")
	rootCmd.Flags().BoolS("H", "H", false, "Follow symlinks on command line")
	rootCmd.Flags().BoolS("L", "L", false, "Follow all symlinks")
	rootCmd.Flags().BoolS("P", "P", false, "Do not follow symlinks")
	rootCmd.Flags().BoolS("X", "X", false, "Permit find to be safely used with xargs")
	rootCmd.Flags().BoolS("d", "d", false, "Depth-first traversal")
	rootCmd.Flags().BoolS("s", "s", false, "Sort entries lexicographically")
	rootCmd.Flags().BoolS("x", "x", false, "Do not descend into directories on different filesystem")

	rootCmd.Flags().StringS("f", "f", "", "Specify path")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}
