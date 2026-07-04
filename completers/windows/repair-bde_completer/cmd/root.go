package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "repair-bde",
	Short: "repair a BitLocker-encrypted drive that cannot be unlocked",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/repair-bde",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
