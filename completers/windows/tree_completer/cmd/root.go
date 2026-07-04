package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tree",
	Short: "graphically display the directory structure of a path or of the disk in a drive",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "use text characters instead of graphic characters to show lines")
	rootCmd.Flags().BoolP("f", "f", false, "display the names of the files in each directory")

	carapace.Gen(rootCmd).PositionalCompletion(carapace.ActionDirectories())
}
