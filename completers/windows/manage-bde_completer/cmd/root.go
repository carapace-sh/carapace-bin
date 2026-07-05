package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "manage-bde",
	Short: "manage BitLocker Drive Encryption",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/manage-bde",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
