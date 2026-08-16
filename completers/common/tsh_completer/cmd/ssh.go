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

	sshCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	sshCmd.Flags().Bool("disable-access-request", false, "Disable automatic resource access requests (DEPRECATED: use --request-mode=off).")
	sshCmd.Flags().StringP("dynamic-forward", "D", "", "Forward localhost connections to remote server using SOCKS5.")
	sshCmd.Flags().BoolP("fork-after-authentication", "f", false, "Run in background after authentication is complete.")
	sshCmd.Flags().StringP("forward", "L", "", "Forward localhost connections to remote server.")
	sshCmd.Flags().BoolP("forward-agent", "A", false, "Forward agent to target node.")
	sshCmd.Flags().String("invite", "", "A comma separated list of people to mark as invited for the session.")
	sshCmd.Flags().Bool("local", false, "Execute command on localhost after connecting to SSH node.")
	sshCmd.Flags().String("log-dir", "", "Directory to log separated command output, when executing on multiple nodes. If set, output from each node will also be labeled in the terminal.")
	sshCmd.Flags().Bool("no-disable-access-request", false, "Disable automatic resource access requests (DEPRECATED: use --request-mode=off).")
	sshCmd.Flags().Bool("no-fork-after-authentication", false, "Run in background after authentication is complete.")
	sshCmd.Flags().Bool("no-forward-agent", false, "Forward agent to target node.")
	sshCmd.Flags().Bool("no-local", false, "Execute command on localhost after connecting to SSH node.")
	sshCmd.Flags().Bool("no-no-remote-exec", false, "Don't execute remote command, useful for port forwarding.")
	sshCmd.Flags().Bool("no-no-resume", false, "Disable SSH connection resumption.")
	sshCmd.Flags().Bool("no-participant-req", false, "Displays a verbose list of required participants in a moderated session.")
	sshCmd.Flags().Bool("no-relogin", false, "Permit performing an authentication attempt on a failed command.")
	sshCmd.Flags().BoolP("no-remote-exec", "N", false, "Don't execute remote command, useful for port forwarding.")
	sshCmd.Flags().Bool("no-resume", false, "Disable SSH connection resumption.")
	sshCmd.Flags().Bool("no-tty", false, "Allocate TTY.")
	sshCmd.Flags().Bool("no-x11-trusted", false, "Requests trusted (insecure) X11 forwarding for this session. This can make your local machine vulnerable to attacks, use with caution.")
	sshCmd.Flags().Bool("no-x11-untrusted", false, "Requests untrusted (secure) X11 forwarding for this session.")
	sshCmd.Flags().StringP("option", "o", "", "OpenSSH options in the format used in the configuration file.")
	sshCmd.Flags().Bool("participant-req", false, "Displays a verbose list of required participants in a moderated session.")
	sshCmd.Flags().StringP("port", "p", "", "SSH port on a remote host.")
	sshCmd.Flags().String("reason", "", "The purpose of the session.")
	sshCmd.Flags().Bool("relogin", true, "Permit performing an authentication attempt on a failed command.")
	sshCmd.Flags().StringP("remote-forward", "R", "", "Forward remote connections to localhost.")
	sshCmd.Flags().String("request-mode", "resource", "Type of automatic access request to make (off, resource, role).")
	sshCmd.Flags().String("request-reason", "", "Reason for requesting access.")
	sshCmd.Flags().BoolP("tty", "t", false, "Allocate TTY.")
	sshCmd.Flags().BoolP("x11-trusted", "Y", false, "Requests trusted (insecure) X11 forwarding for this session. This can make your local machine vulnerable to attacks, use with caution.")
	sshCmd.Flags().BoolP("x11-untrusted", "X", false, "Requests untrusted (secure) X11 forwarding for this session.")
	sshCmd.Flags().String("x11-untrusted-timeout", "10m", "Sets a timeout for untrusted X11 forwarding, after which the client will reject any forwarding requests from the server.")
	sshCmd.Flag("no-disable-access-request").Hidden = true
	sshCmd.Flag("no-fork-after-authentication").Hidden = true
	sshCmd.Flag("no-forward-agent").Hidden = true
	sshCmd.Flag("no-local").Hidden = true
	sshCmd.Flag("no-no-remote-exec").Hidden = true
	sshCmd.Flag("no-no-resume").Hidden = true
	sshCmd.Flag("no-participant-req").Hidden = true
	sshCmd.Flag("no-relogin").Hidden = true
	sshCmd.Flag("no-tty").Hidden = true
	sshCmd.Flag("no-x11-trusted").Hidden = true
	sshCmd.Flag("no-x11-untrusted").Hidden = true
	rootCmd.AddCommand(sshCmd)
}
