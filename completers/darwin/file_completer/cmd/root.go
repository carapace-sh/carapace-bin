package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file",
	Short: "determine file type",
	Long:  "https://keith.github.io/xcode-manpages/file.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("D", "D", false, "Show compiled magic file")
	rootCmd.Flags().BoolS("I", "I", false, "Show MIME type and encoding")
	rootCmd.Flags().BoolS("L", "L", false, "Follow symlinks")
	rootCmd.Flags().BoolS("N", "N", false, "Do not pad file name")
	rootCmd.Flags().BoolS("b", "b", false, "Do not prepend file name")
	rootCmd.Flags().BoolS("c", "c", false, "Show magic file checks")
	rootCmd.Flags().BoolS("d", "d", false, "Show debugging output")
	rootCmd.Flags().BoolS("h", "h", false, "Do not follow symlinks")
	rootCmd.Flags().BoolS("i", "i", false, "Show MIME type and encoding")
	rootCmd.Flags().BoolS("k", "k", false, "Do not stop at first match")
	rootCmd.Flags().BoolS("n", "n", false, "Flush stdout after each file")
	rootCmd.Flags().BoolS("p", "p", false, "Preserve file access time")
	rootCmd.Flags().BoolS("r", "r", false, "Do not translate unprintable characters")
	rootCmd.Flags().BoolS("s", "s", false, "Read special files")
	rootCmd.Flags().BoolS("v", "v", false, "Print version")
	rootCmd.Flags().BoolS("z", "z", false, "Examine compressed files")

	rootCmd.Flags().StringS("M", "M", "", "Use compiled magic files")
	rootCmd.Flags().StringS("f", "f", "", "Read file names from namefile")
	rootCmd.Flags().StringS("m", "m", "", "Use magic files")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"M": carapace.ActionFiles(),
		"f": carapace.ActionFiles(),
		"m": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
