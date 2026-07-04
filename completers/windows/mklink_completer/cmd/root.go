package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mklink",
	Short: "create a symbolic or hard link",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/mklink",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("d", "d", false, "create a directory symbolic link")
	rootCmd.Flags().BoolP("h", "h", false, "create a hard link instead of a symbolic link")
	rootCmd.Flags().BoolP("j", "j", false, "create a directory junction")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
