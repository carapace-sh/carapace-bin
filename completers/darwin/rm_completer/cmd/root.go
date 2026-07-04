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

	rootCmd.Flags().BoolS("I", "I", false, "Request confirmation once if more than three files are being removed")
	rootCmd.Flags().BoolS("R", "R", false, "Attempt to remove the file hierarchy rooted in each file argument")
	rootCmd.Flags().BoolS("W", "W", false, "Attempt to undelete the named files")
	rootCmd.Flags().BoolP("directory", "d", false, "Attempt to remove directories as well as other types of files")
	rootCmd.Flags().BoolP("force", "f", false, "Attempt to remove the files without prompting for confirmation")
	rootCmd.Flags().BoolP("interactive", "i", false, "Request confirmation before attempting to remove each file")
	rootCmd.Flags().BoolP("one-file-system", "x", false, "When removing a hierarchy, do not cross mount points")
	rootCmd.Flags().BoolP("recursive", "r", false, "Equivalent to -R")
	rootCmd.Flags().BoolP("verbose", "v", false, "Be verbose when deleting files, showing them as they are removed")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
