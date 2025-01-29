package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machine_inspectCmd = &cobra.Command{
	Use:   "inspect [options] [MACHINE...]",
	Short: "Inspect an existing machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_inspectCmd).Standalone()

	machine_inspectCmd.Flags().String("format", "", "Format volume output using JSON or a Go template")
	machineCmd.AddCommand(machine_inspectCmd)
}
