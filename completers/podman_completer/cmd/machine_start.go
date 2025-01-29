package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machine_startCmd = &cobra.Command{
	Use:   "start [options] [MACHINE]",
	Short: "Start an existing machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_startCmd).Standalone()

	machine_startCmd.Flags().Bool("no-info", false, "Suppress informational tips")
	machine_startCmd.Flags().BoolP("quiet", "q", false, "Suppress machine starting status output")
	machineCmd.AddCommand(machine_startCmd)
}
