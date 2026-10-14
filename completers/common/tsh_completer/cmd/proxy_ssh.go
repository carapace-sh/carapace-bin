package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proxy_sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Start local TLS proxy for ssh connections when using Teleport in single-port mode.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxy_sshCmd).Standalone()

	proxy_sshCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	proxy_sshCmd.Flags().Bool("no-no-resume", false, "Disable SSH connection resumption.")
	proxy_sshCmd.Flags().Bool("no-relogin", false, "Permit performing an authentication attempt on a failed command.")
	proxy_sshCmd.Flags().Bool("no-resume", false, "Disable SSH connection resumption.")
	proxy_sshCmd.Flags().Bool("relogin", true, "Permit performing an authentication attempt on a failed command.")
	proxy_sshCmd.Flag("no-no-resume").Hidden = true
	proxy_sshCmd.Flag("no-relogin").Hidden = true
	proxyCmd.AddCommand(proxy_sshCmd)
}
