package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "driverquery",
	Short: "display a list of all installed device drivers and their properties",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/driverquery",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("nh", "nh", false, "skip header row")
	rootCmd.Flags().BoolP("si", "si", false, "display signed drivers")
	rootCmd.Flags().BoolP("v", "v", false, "display verbose information")
}
