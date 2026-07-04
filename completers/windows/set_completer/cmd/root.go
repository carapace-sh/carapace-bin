package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "set",
	Short: "display, set, or remove cmd.exe environment variables",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "evaluate numerical expressions")
	rootCmd.Flags().BoolP("p", "p", false, "set variable to a line of input")
}
