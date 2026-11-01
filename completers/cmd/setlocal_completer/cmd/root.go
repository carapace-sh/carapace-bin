package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "setlocal",
	Short: "start localization of environment variables in a batch file",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/setlocal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("enabledelayedexpansion", false, "enable delayed environment variable expansion")
	rootCmd.Flags().Bool("enableextensions", false, "enable command processor extensions")
}
