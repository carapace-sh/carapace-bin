package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bdehdcfg",
	Short: "prepare a drive for BitLocker Drive Encryption",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/bdehdcfg",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
