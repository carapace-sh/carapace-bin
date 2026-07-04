package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "start a separate window to run a specified program or command",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/start",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("b", "b", false, "start application without creating a new window")
	rootCmd.Flags().BoolP("min", "min", false, "start window minimized")
	rootCmd.Flags().BoolP("max", "max", false, "start window maximized")
	rootCmd.Flags().BoolP("wait", "wait", false, "start application and wait for it to terminate")
}
