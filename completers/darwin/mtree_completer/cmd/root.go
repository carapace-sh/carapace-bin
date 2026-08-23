package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mtree",
	Short: "compare file hierarchies",
	Long:  "https://keith.github.io/xcode-manpages/mtree.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("D", "D", false, "Do not set XATTR_SHOWCOMPRESSION flag")
	rootCmd.Flags().BoolS("L", "L", false, "Follow all symbolic links")
	rootCmd.Flags().BoolS("P", "P", false, "Do not follow symbolic links")
	rootCmd.Flags().BoolS("S", "S", false, "Skip calculating digest of extended attributes")
	rootCmd.Flags().BoolS("U", "U", false, "Modify owner, group, permissions, and modification time")
	rootCmd.Flags().BoolS("c", "c", false, "Print a specification for the file hierarchy")
	rootCmd.Flags().BoolS("d", "d", false, "Ignore everything except directory type files")
	rootCmd.Flags().BoolS("e", "e", false, "Do not complain about files not in the spec")
	rootCmd.Flags().BoolS("i", "i", false, "Indent the output 4 spaces per directory level")
	rootCmd.Flags().BoolS("n", "n", false, "Do not emit pathname comments")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode")
	rootCmd.Flags().BoolS("r", "r", false, "Remove files not described in the specification")
	rootCmd.Flags().BoolS("u", "u", false, "Same as -U, but return status 2 if changes fail")
	rootCmd.Flags().BoolS("w", "w", false, "Make some error conditions non-fatal warnings")
	rootCmd.Flags().BoolS("x", "x", false, "Do not descend below mount points")

	rootCmd.Flags().StringS("K", "K", "", "Add the specified keywords to the current set")
	rootCmd.Flags().StringS("f", "f", "", "Read the specification from file")
	rootCmd.Flags().StringS("k", "k", "", "Use the specified keywords")
	rootCmd.Flags().StringS("p", "p", "", "Use the specified path as the root of the file hierarchy")
	rootCmd.Flags().StringS("s", "s", "", "Use the specified seed for the hash")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionFiles(),
		"p": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}