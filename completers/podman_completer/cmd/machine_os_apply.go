package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machine_os_applyCmd = &cobra.Command{
	Use:   "apply [options] IMAGE [NAME]",
	Short: "Apply an OCI image to a Podman Machine's OS",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_os_applyCmd).Standalone()

	machine_os_applyCmd.Flags().Bool("restart", false, "Restart VM to apply changes")
	machine_osCmd.AddCommand(machine_os_applyCmd)
}
