package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "time",
	Short: "display or set the system time",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/time",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("t", "t", false, "display the current time without prompting for a new time")
}
