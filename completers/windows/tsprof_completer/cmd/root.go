package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tsprof",
	Short: "copy user configuration from one user to another on a terminal server",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/tsprof",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
