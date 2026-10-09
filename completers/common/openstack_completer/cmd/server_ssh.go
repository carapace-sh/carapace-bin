package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH to server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_sshCmd).Standalone()

	server_sshCmd.Flags().BoolS("4", "4", false, "Use only IPv4 addresses")
	server_sshCmd.Flags().BoolS("6", "6", false, "Use only IPv6 addresses")
	server_sshCmd.Flags().String("address-type", "", "Use other IP address (public, private, etc)")
	server_sshCmd.Flags().StringP("identity", "i", "", "==SUPPRESS==")
	server_sshCmd.Flags().StringP("login", "l", "", "==SUPPRESS==")
	server_sshCmd.Flags().StringP("option", "o", "", "==SUPPRESS==")
	server_sshCmd.Flags().StringP("port", "p", "", "==SUPPRESS==")
	server_sshCmd.Flags().Bool("private", false, "Use private IP address")
	server_sshCmd.Flags().Bool("public", false, "Use public IP address")
	server_sshCmd.Flags().BoolS("v", "v", false, "==SUPPRESS==")
	serverCmd.AddCommand(server_sshCmd)
}
