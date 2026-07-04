package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "robocopy",
	Short: "robust file copy for Windows",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/robocopy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "copy only files with the archive attribute set")
	rootCmd.Flags().BoolP("b", "b", false, "copy files in backup mode (unrestricted access)")
	rootCmd.Flags().BoolP("c", "c", false, "continue copying even if errors occur")
	rootCmd.Flags().BoolP("e", "e", false, "copy subdirectories, including empty ones")
	rootCmd.Flags().BoolP("j", "j", false, "copy using unbuffered I/O (recommended for large files)")
	rootCmd.Flags().Bool("mir", false, "mirror a directory tree (equivalent to /e plus /purge)")
	rootCmd.Flags().Bool("move", false, "move files and directories (delete from source after copying)")
	rootCmd.Flags().Bool("ndl", false, "no directory list (don't log directory names)")
	rootCmd.Flags().Bool("nfl", false, "no file list (don't log file names)")
	rootCmd.Flags().Bool("njh", false, "no job header")
	rootCmd.Flags().Bool("njs", false, "no job summary")
	rootCmd.Flags().Bool("np", false, "no progress indicator")
	rootCmd.Flags().Bool("purge", false, "delete dest files/dirs that no longer exist in source")
	rootCmd.Flags().BoolP("r", "r", false, "number of retries on failed copies")
	rootCmd.Flags().BoolP("s", "s", false, "copy subdirectories, but not empty ones")
	rootCmd.Flags().BoolP("w", "w", false, "wait time between retries in seconds")
	rootCmd.Flags().Bool("xo", false, "exclude older files")
	rootCmd.Flags().Bool("xx", false, "exclude extra files and directories")
	rootCmd.Flags().Bool("zf", false, "use Zip mode (journaling for restartability)")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		carapace.ActionDirectories(),
	)
}
