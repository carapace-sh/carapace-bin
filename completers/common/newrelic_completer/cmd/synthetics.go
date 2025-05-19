package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var syntheticsCmd = &cobra.Command{
	Use:   "synthetics",
	Short: "Interact with New Relic Synthetics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syntheticsCmd).Standalone()
	rootCmd.AddCommand(syntheticsCmd)
}
