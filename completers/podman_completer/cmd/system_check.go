package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_checkCmd = &cobra.Command{
	Use:   "check [options]",
	Short: "Check storage consistency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_checkCmd).Standalone()

	system_checkCmd.Flags().BoolP("force", "f", false, "Remove inconsistent images and containers")
	system_checkCmd.Flags().StringP("max", "m", "", "Maximum allowed age of unreferenced layers")
	system_checkCmd.Flags().BoolP("quick", "q", false, "Skip time-consuming checks. The default is to include time-consuming checks")
	system_checkCmd.Flags().BoolP("repair", "r", false, "Remove inconsistent images")
	systemCmd.AddCommand(system_checkCmd)
}
