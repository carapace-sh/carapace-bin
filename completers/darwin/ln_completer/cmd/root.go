package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ln",
	Short: "make links between files",
	Long:  "https://keith.github.io/xcode-manpages/ln.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("F", "F", false, "Force existing destination path to be removed to allow the link to be created")
	rootCmd.Flags().BoolS("L", "L", false, "Create a symbolic link to the target file or directory")
	rootCmd.Flags().BoolS("P", "P", false, "Create a hard link to the target file or directory")
	rootCmd.Flags().BoolS("f", "f", false, "If the target file already exists, then unlink it so that the link may occur")
	rootCmd.Flags().BoolS("h", "h", false, "If the target_file or target_dir is a symbolic link to a directory, do not follow it")
	rootCmd.Flags().BoolS("i", "i", false, "Cause ln to write a prompt to standard error if the target file exists")
	rootCmd.Flags().BoolS("n", "n", false, "Same as -h, but if the target is a symbolic link to a directory, do not follow it")
	rootCmd.Flags().BoolS("s", "s", false, "Create a symbolic link")
	rootCmd.Flags().BoolS("v", "v", false, "Cause ln to be verbose, showing files as they are processed")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
