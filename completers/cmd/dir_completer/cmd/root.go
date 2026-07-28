package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dir",
	Short: "display a list of files and subdirectories in a directory",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/dir",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("4", "4", false, "display four-digit year")
	rootCmd.Flags().BoolP("a", "a", false, "display files with specified attributes")
	rootCmd.Flags().BoolP("b", "b", false, "use bare format (no heading information or summary)")
	rootCmd.Flags().BoolP("c", "c", false, "display the thousand separator in file sizes")
	rootCmd.Flags().BoolP("d", "d", false, "display in same format as /w but sorted by column")
	rootCmd.Flags().BoolP("l", "l", false, "display lowercase")
	rootCmd.Flags().BoolP("n", "n", false, "display in new long list format")
	rootCmd.Flags().BoolP("o", "o", false, "sort by specified order")
	rootCmd.Flags().BoolP("p", "p", false, "pause after each screenful of information")
	rootCmd.Flags().BoolP("q", "q", false, "display the owner of the file")
	rootCmd.Flags().BoolP("r", "r", false, "display alternate data streams of the file")
	rootCmd.Flags().BoolP("s", "s", false, "display files in specified directory and all subdirectories")
	rootCmd.Flags().BoolP("t", "t", false, "control which time field to display")
	rootCmd.Flags().BoolP("w", "w", false, "display in wide list format")
	rootCmd.Flags().BoolP("x", "x", false, "display short names for non-8.3 files")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}
