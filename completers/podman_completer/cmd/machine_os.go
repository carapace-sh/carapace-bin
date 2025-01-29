package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machine_osCmd = &cobra.Command{
	Use:   "os",
	Short: "Manage a Podman virtual machine's OS",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_osCmd).Standalone()

	machineCmd.AddCommand(machine_osCmd)
}
