package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tunnel_6rdCmd = &cobra.Command{
	Use:   "6rd",
	Short: "6rd (IPv6 Rapid Deployment) configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tunnel_6rdCmd).Standalone()

	tunnelCmd.AddCommand(tunnel_6rdCmd)
}
