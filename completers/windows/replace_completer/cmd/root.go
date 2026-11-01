package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "replace",
	Short: "replace files in the destination directory with files of the same name in the source directory",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/replace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "add new files to destination directory")
	rootCmd.Flags().BoolP("p", "p", false, "prompt before replacing or adding a file")
	rootCmd.Flags().BoolP("r", "r", false, "replace read-only files")
	rootCmd.Flags().BoolP("s", "s", false, "replace files in all subdirectories")
	rootCmd.Flags().BoolP("u", "u", false, "replace (update) only files that are older")
	rootCmd.Flags().BoolP("w", "w", false, "wait for keypress before processing")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionDirectories(),
	)
}
