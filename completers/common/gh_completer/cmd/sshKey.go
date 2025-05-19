package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sshKeyCmd = &cobra.Command{
	Use:   "ssh-key <command>",
	Short: "Manage SSH keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKeyCmd).Standalone()

	rootCmd.AddCommand(sshKeyCmd)
}
