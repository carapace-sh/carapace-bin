package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_sshKeyCmd = &cobra.Command{
	Use:   "ssh-key",
	Short: "Display commands to manage SSH keys on your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_sshKeyCmd).Standalone()
	computeCmd.AddCommand(compute_sshKeyCmd)
}
