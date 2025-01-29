package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machine_sshCmd = &cobra.Command{
	Use:   "ssh [options] [NAME] [COMMAND [ARG ...]]",
	Short: "SSH into an existing machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_sshCmd).Standalone()

	machine_sshCmd.Flags().String("username", "", "Username to use when ssh-ing into the VM.")
	machineCmd.AddCommand(machine_sshCmd)
}
