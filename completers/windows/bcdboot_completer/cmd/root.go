package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bcdboot",
	Short: "set up system partition or repair Boot Configuration Data",
	Long:  "https://learn.microsoft.com/en-us/windows-hardware/manufacture/desktop/bcdboot-command-line-options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
