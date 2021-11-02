package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var compute_sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Access a Droplet using SSH",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_sshCmd).Standalone()
	compute_sshCmd.Flags().Bool("ssh-agent-forwarding", false, "Enable SSH agent forwarding")
	compute_sshCmd.Flags().String("ssh-command", "", "Command to execute on Droplet")
	compute_sshCmd.Flags().String("ssh-key-path", "/home/rsteube/.ssh/id_rsa", "Path to SSH private key")
	compute_sshCmd.Flags().Int("ssh-port", 22, "The remote port sshd is running on")
	compute_sshCmd.Flags().Bool("ssh-private-ip", false, "SSH to Droplet's private IP address")
	compute_sshCmd.Flags().String("ssh-user", "root", "SSH user for connection")
	computeCmd.AddCommand(compute_sshCmd)

	carapace.Gen(compute_sshCmd).FlagCompletion(carapace.ActionMap{
		"ssh-key-path": carapace.ActionValues(),
		"ssh-user":     os.ActionUsers(),
	})
}
