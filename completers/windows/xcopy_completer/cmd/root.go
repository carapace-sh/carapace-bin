package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xcopy",
	Short: "copy files and directories, including subdirectories",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/xcopy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "copy only files with the archive attribute set")
	rootCmd.Flags().BoolP("c", "c", false, "continue copying even if errors occur")
	rootCmd.Flags().BoolP("d", "d", false, "copy only files whose source time is newer than destination")
	rootCmd.Flags().BoolP("e", "e", false, "copy all subdirectories, even if they are empty")
	rootCmd.Flags().BoolP("h", "h", false, "copy hidden and system files")
	rootCmd.Flags().BoolP("i", "i", false, "if destination does not exist, assume it is a directory")
	rootCmd.Flags().BoolP("k", "k", false, "copy attributes, read-only by default")
	rootCmd.Flags().BoolP("q", "q", false, "quiet mode, do not display file names while copying")
	rootCmd.Flags().BoolP("r", "r", false, "overwrite read-only files")
	rootCmd.Flags().BoolP("s", "s", false, "copy directories and subdirectories except empty ones")
	rootCmd.Flags().BoolP("t", "t", false, "copy directory structure only, not files")
	rootCmd.Flags().BoolP("u", "u", false, "copy only files that already exist in destination")
	rootCmd.Flags().BoolP("w", "w", false, "prompt to press a key before copying")
	rootCmd.Flags().BoolP("x", "x", false, "copy file audit settings and ACLs (implies /o)")
	rootCmd.Flags().BoolP("y", "y", false, "suppress prompting to confirm overwriting")
	rootCmd.Flags().BoolP("z", "z", false, "copy files in restartable mode")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
