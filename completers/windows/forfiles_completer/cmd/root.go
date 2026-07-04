package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "forfiles",
	Short: "select files by date or pattern for batch processing",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/forfiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("p", "p", false, "specify the search path")
	rootCmd.Flags().BoolP("s", "s", false, "search subdirectories")
}
