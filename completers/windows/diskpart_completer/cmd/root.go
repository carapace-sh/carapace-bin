package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "diskpart",
	Short: "disk partitioning and volume management scripting tool",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/diskpart",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
