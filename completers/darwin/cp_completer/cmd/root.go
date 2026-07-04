package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cp",
	Short: "copy files",
	Long:  "https://keith.github.io/xcode-manpages/cp.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("archive", "a", false, "Same as -pPR options")
	rootCmd.Flags().BoolP("clipboard", "c", false, "Copy the file data to the pasteboard")
	rootCmd.Flags().BoolP("follow-all-symlinks", "L", false, "Follow all symbolic links encountered during traversal")
	rootCmd.Flags().BoolP("follow-symlinks", "H", false, "Follow symbolic links encountered as arguments")
	rootCmd.Flags().BoolP("force", "f", false, "If the destination file cannot be opened, remove it and create a new file")
	rootCmd.Flags().BoolP("interactive", "i", false, "Cause cp to write a prompt to the standard error output before copying a file")
	rootCmd.Flags().BoolP("no-clobber", "n", false, "Do not overwrite an existing file")
	rootCmd.Flags().BoolP("no-dereference", "P", false, "Make the copy a symbolic link to the original")
	rootCmd.Flags().BoolP("no-preserve-extended-attributes", "N", false, "Do not preserve extended attributes and ACLs")
	rootCmd.Flags().BoolP("no-traverse", "X", false, "Do not copy directories that are on a different filesystem")
	rootCmd.Flags().BoolP("preserve", "p", false, "Preserve the following attributes of each source file in the copy")
	rootCmd.Flags().BoolP("recursive", "R", false, "If the source_file designates a directory, copy the directory and the entire subtree")
	rootCmd.Flags().BoolP("verbose", "v", false, "Cause cp to be verbose, showing files as they are copied")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
