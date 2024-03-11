package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Interact with New Relic Synthetics monitors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_monitorCmd).Standalone()
	syntheticsCmd.AddCommand(synthetics_monitorCmd)
}
