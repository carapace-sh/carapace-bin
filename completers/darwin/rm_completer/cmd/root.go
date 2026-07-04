package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove directory entries",
	Long:  "https://keith.github.io/xcode-manpages/rm.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("directory", "d", false, "Attempt to remove directories and files in their hierarchies, including empty directories")
	rootCmd.Flags().BoolP("force", "f", false, "Attempt to remove the files without prompting for confirmation")
	rootCmd.Flags().BoolP("interactive", "i", false, "Request confirmation before attempting to remove each file")
	rootCmd.Flags().BoolP("overwrite", "P", false, "Overwrite regular files before deleting them")
	rootCmd.Flags().BoolP("recursive", "R", false, "Attempt to remove the file hierarchy rooted in each file argument")
	rootCmd.Flags().BoolP("recursive-alt", "r", false, "Attempt to remove the file hierarchy rooted in each file argument")
	rootCmd.Flags().BoolP("verbose", "v", false, "Be verbose when carrying out actions")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
