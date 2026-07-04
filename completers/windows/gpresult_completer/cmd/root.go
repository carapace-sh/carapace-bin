package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gpresult",
	Short: "display group policy information for a user or computer",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/gpresult",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("r", "r", false, "display RSoP summary data")
	rootCmd.Flags().BoolP("v", "v", false, "display verbose policy information")
	rootCmd.Flags().BoolP("z", "z", false, "display all available information")
}
