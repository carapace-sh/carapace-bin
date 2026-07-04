package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "shut down or restart a local or remote computer",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/shutdown",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "abort a system shutdown")
	rootCmd.Flags().BoolP("f", "f", false, "force running applications to close")
	rootCmd.Flags().BoolP("h", "h", false, "hibernate the local computer")
	rootCmd.Flags().BoolP("l", "l", false, "log off the current user")
	rootCmd.Flags().BoolP("p", "p", false, "turn off the local computer with no timeout or warning")
	rootCmd.Flags().BoolP("r", "r", false, "restart the computer")
	rootCmd.Flags().BoolP("s", "s", false, "shut down the computer")
	rootCmd.Flags().BoolP("t", "t", false, "set shutdown timeout in seconds")
}
