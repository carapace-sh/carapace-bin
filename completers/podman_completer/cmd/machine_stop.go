package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machine_stopCmd = &cobra.Command{
	Use:   "stop [MACHINE]",
	Short: "Stop an existing machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_stopCmd).Standalone()

	machineCmd.AddCommand(machine_stopCmd)
}
