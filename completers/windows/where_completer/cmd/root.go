package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "where",
	Short: "display the location of files that match the search pattern",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/where",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("f", "f", false, "show the results of the pattern in both upper and lower case")
	rootCmd.Flags().BoolP("q", "q", false, "quiet mode, return only exit code")
	rootCmd.Flags().BoolP("r", "r", false, "recursive search, display files in specified directory and subdirectories")
	rootCmd.Flags().BoolP("t", "t", false, "display the file size, timestamp, and date")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
