package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "del",
	Short: "delete one or more files",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/del",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("f", "f", false, "force deletion of read-only files")
	rootCmd.Flags().BoolP("p", "p", false, "prompt for confirmation before deleting each file")
	rootCmd.Flags().BoolP("q", "q", false, "quiet mode, do not prompt for confirmation")
	rootCmd.Flags().BoolP("s", "s", false, "delete specified files from all subdirectories")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
