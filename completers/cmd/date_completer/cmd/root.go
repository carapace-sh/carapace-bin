package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "date",
	Short: "display or set the current date",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/date",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("t", "t", false, "display the current date without prompting for a new date")
}
