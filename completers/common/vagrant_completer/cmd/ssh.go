package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "connects to machine via SSH",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshCmd).Standalone()

	sshCmd.Flags().StringP("command", "c", "", "Execute an SSH command directly")
	sshCmd.Flags().Bool("no-tty", false, "Disables tty when executing an ssh command (defaults to true)")
	sshCmd.Flags().BoolP("plain", "p", false, "Plain mode, leaves authentication up to user")
	sshCmd.Flags().BoolP("tty", "t", false, "Enables tty when executing an ssh command (defaults to true)")
	rootCmd.AddCommand(sshCmd)

	carapace.Gen(sshCmd).PositionalCompletion(
		vagrant.ActionMachines(),
	)
}
