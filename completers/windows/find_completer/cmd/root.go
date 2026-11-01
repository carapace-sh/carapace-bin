package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "find",
	Short: "search for a text string in a file or files",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/find",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("c", "c", false, "display only the count of lines containing the string")
	rootCmd.Flags().BoolP("i", "i", false, "case-insensitive search")
	rootCmd.Flags().BoolP("n", "n", false, "display line numbers with each line")
	rootCmd.Flags().BoolP("v", "v", false, "display all lines not containing the string")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
