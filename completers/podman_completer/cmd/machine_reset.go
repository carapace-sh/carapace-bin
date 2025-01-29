package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machine_resetCmd = &cobra.Command{
	Use:   "reset [options]",
	Short: "Remove all machines",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_resetCmd).Standalone()

	machine_resetCmd.Flags().BoolP("force", "f", false, "Do not prompt before reset")
	machineCmd.AddCommand(machine_resetCmd)
}
