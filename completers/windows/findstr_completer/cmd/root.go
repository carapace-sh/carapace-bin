package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "findstr",
	Short: "search for text in files",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/findstr",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("b", "b", false, "match pattern if at beginning of a line")
	rootCmd.Flags().BoolP("e", "e", false, "match pattern if at end of a line")
	rootCmd.Flags().BoolP("i", "i", false, "case-insensitive search")
	rootCmd.Flags().BoolP("n", "n", false, "print line number with each line")
	rootCmd.Flags().BoolP("o", "o", false, "print character offset before each matching line")
	rootCmd.Flags().BoolP("p", "p", false, "skip files with non-printable characters")
	rootCmd.Flags().BoolP("r", "r", false, "use search strings as regular expressions")
	rootCmd.Flags().BoolP("s", "s", false, "search current directory and all subdirectories")
	rootCmd.Flags().BoolP("v", "v", false, "print lines that do not contain a match")
	rootCmd.Flags().BoolP("x", "x", false, "print lines that match exactly")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
