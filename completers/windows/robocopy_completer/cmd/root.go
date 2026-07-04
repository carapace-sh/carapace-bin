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
	rootCmd.Flags().BoolP("mir", "mir", false, "mirror a directory tree (equivalent to /e plus /purge)")
	rootCmd.Flags().BoolP("move", "move", false, "move files and directories (delete from source after copying)")
	rootCmd.Flags().BoolP("njh", "njh", false, "no job header")
	rootCmd.Flags().BoolP("njs", "njs", false, "no job summary")
	rootCmd.Flags().BoolP("np", "np", false, "no progress indicator")
	rootCmd.Flags().BoolP("nfl", "nfl", false, "no file list (don't log file names)")
	rootCmd.Flags().BoolP("ndl", "ndl", false, "no directory list (don't log directory names)")
	rootCmd.Flags().BoolP("purge", "purge", false, "delete dest files/dirs that no longer exist in source")
	rootCmd.Flags().BoolP("r:0", "r", false, "retry 0 times on failed copies")
	rootCmd.Flags().BoolP("s", "s", false, "copy subdirectories, but not empty ones")
	rootCmd.Flags().BoolP("w:0", "w", false, "wait 0 seconds between retries")
	rootCmd.Flags().BoolP("xo", "xo", false, "exclude older files")
	rootCmd.Flags().BoolP("xx", "xx", false, "exclude extra files and directories")
	rootCmd.Flags().BoolP("zf", "zf", false, "use Zip mode (journaling for restartability)")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		carapace.ActionDirectories(),
	)
}
