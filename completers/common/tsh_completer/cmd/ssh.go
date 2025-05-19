package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Run shell or execute a command on a remote SSH node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshCmd).Standalone()

	sshCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	sshCmd.Flags().Bool("disable-access-request", false, "Disable automatic resource access requests")
	sshCmd.Flags().StringP("dynamic-forward", "D", "", "Forward localhost connections to remote server using SOCKS5")
	sshCmd.Flags().StringP("forward", "L", "", "Forward localhost connections to remote server")
	sshCmd.Flags().BoolP("forward-agent", "A", false, "Forward agent to target node")
	sshCmd.Flags().Bool("local", false, "Execute command on localhost after connecting to SSH node")
	sshCmd.Flags().Bool("no-disable-access-request", false, "Disable automatic resource access requests")
	sshCmd.Flags().Bool("no-forward-agent", false, "Forward agent to target node")
	sshCmd.Flags().Bool("no-local", false, "Execute command on localhost after connecting to SSH node")
	sshCmd.Flags().Bool("no-no-remote-exec", false, "Don't execute remote command, useful for port forwarding")
	sshCmd.Flags().Bool("no-participant-req", false, "Displays a verbose list of required participants in a moderated session.")
	sshCmd.Flags().BoolP("no-remote-exec", "N", false, "Don't execute remote command, useful for port forwarding")
	sshCmd.Flags().Bool("no-tty", false, "Allocate TTY")
	sshCmd.Flags().Bool("no-x11-trusted", false, "Requests trusted (insecure) X11 forwarding for this session. This can make your local machine vulnerable to attacks, use with caution")
	sshCmd.Flags().Bool("no-x11-untrusted", false, "Requests untrusted (secure) X11 forwarding for this session")
	sshCmd.Flags().StringP("option", "o", "", "OpenSSH options in the format used in the configuration file")
	sshCmd.Flags().Bool("participant-req", false, "Displays a verbose list of required participants in a moderated session.")
	sshCmd.Flags().StringP("port", "p", "", "SSH port on a remote host")
	sshCmd.Flags().String("request-reason", "", "Reason for requesting access")
	sshCmd.Flags().BoolP("tty", "t", false, "Allocate TTY")
	sshCmd.Flags().BoolP("x11-trusted", "Y", false, "Requests trusted (insecure) X11 forwarding for this session. This can make your local machine vulnerable to attacks, use with caution")
	sshCmd.Flags().BoolP("x11-untrusted", "X", false, "Requests untrusted (secure) X11 forwarding for this session")
	sshCmd.Flags().String("x11-untrusted-timeout", "", "Sets a timeout for untrusted X11 forwarding, after which the client will reject any forwarding requests from the server")
	sshCmd.Flag("no-disable-access-request").Hidden = true
	sshCmd.Flag("no-forward-agent").Hidden = true
	sshCmd.Flag("no-local").Hidden = true
	sshCmd.Flag("no-no-remote-exec").Hidden = true
	sshCmd.Flag("no-participant-req").Hidden = true
	sshCmd.Flag("no-tty").Hidden = true
	sshCmd.Flag("no-x11-trusted").Hidden = true
	sshCmd.Flag("no-x11-untrusted").Hidden = true
	rootCmd.AddCommand(sshCmd)
}
