package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stat",
	Short: "display file or file system status",
	Long:  "https://keith.github.io/xcode-manpages/stat.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("F", "F", false, "Display a slash (/) immediately after each pathname that is a directory")
	rootCmd.Flags().BoolS("L", "L", false, "Use stat(2) instead of lstat(2)")
	rootCmd.Flags().StringS("f", "f", "", "Display information using the specified format")
	rootCmd.Flags().BoolS("l", "l", false, "Display output in ls -lT format")
	rootCmd.Flags().BoolS("n", "n", false, "Do not force a newline to be printed at the end of each operand")
	rootCmd.Flags().BoolS("q", "q", false, "Fail silently if any of the files given do not exist")
	rootCmd.Flags().BoolS("r", "r", false, "Display raw information (no special interpretation)")
	rootCmd.Flags().BoolS("s", "s", false, "Display file status flags in symbolic form")
	rootCmd.Flags().StringS("t", "t", "", "Display the time in the specified timefmt")
	rootCmd.Flags().BoolS("x", "x", false, "Display information in a more verbose, XML-like format")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
