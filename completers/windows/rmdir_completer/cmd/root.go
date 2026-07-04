package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rmdir",
	Short: "remove a directory",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/rmdir",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("s", "s", false, "remove all files and subdirectories in the specified directory")
	rootCmd.Flags().BoolP("q", "q", false, "quiet mode, do not prompt for confirmation")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}
