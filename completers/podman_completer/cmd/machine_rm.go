package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machine_rmCmd = &cobra.Command{
	Use:   "rm [options] [MACHINE]",
	Short: "Remove an existing machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_rmCmd).Standalone()

	machine_rmCmd.Flags().BoolP("force", "f", false, "Stop and do not prompt before rming")
	machine_rmCmd.Flags().Bool("save-ignition", false, "Do not delete ignition file")
	machine_rmCmd.Flags().Bool("save-image", false, "Do not delete the image file")
	machineCmd.AddCommand(machine_rmCmd)
}
