package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sfc",
	Short: "scan and replace corrupted system files",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/sfc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("scannow", "scannow", false, "scan all protected system files immediately")
	rootCmd.Flags().BoolP("scanonce", "scanonce", false, "scan all protected system files at next boot")
	rootCmd.Flags().BoolP("scanboot", "scanboot", false, "scan all protected system files at every boot")
	rootCmd.Flags().BoolP("revert", "revert", false, "revert to default scanning behavior")
}
