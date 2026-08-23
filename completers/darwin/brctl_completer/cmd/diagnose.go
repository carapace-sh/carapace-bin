package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "diagnose and collect logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagnoseCmd).Standalone()
	rootCmd.AddCommand(diagnoseCmd)

	diagnoseCmd.Flags().Bool("collect-mobile-documents", false, "collect mobile documents")
	diagnoseCmd.Flags().String("name", "", "change the device name")
	diagnoseCmd.Flags().Bool("sysdiagnose", false, "do not collect what's already part of sysdiagnose")

	carapace.Gen(diagnoseCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
