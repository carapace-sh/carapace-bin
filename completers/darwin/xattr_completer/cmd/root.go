package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xattr",
	Short: "display and manipulate extended attributes",
	Long:  "https://keith.github.io/xcode-manpages/xattr.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("clear", "c", false, "Remove all attributes")
	rootCmd.Flags().BoolP("delete", "d", false, "Delete the given attribute name")
	rootCmd.Flags().BoolP("help", "h", false, "Display help message")
	rootCmd.Flags().BoolP("hex", "x", false, "Force attribute value in hexadecimal representation")
	rootCmd.Flags().BoolP("long", "l", false, "Display attribute names and values")
	rootCmd.Flags().BoolP("print", "p", false, "Print the value associated with the given attribute name")
	rootCmd.Flags().BoolP("recursive", "r", false, "Act on the entire contents of a directory recursively")
	rootCmd.Flags().BoolP("symlink", "s", false, "Act on the symbolic link itself")
	rootCmd.Flags().BoolP("verbose", "v", false, "Force the file name to be displayed")
	rootCmd.Flags().BoolP("write", "w", false, "Write the given attribute name with the given value")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}