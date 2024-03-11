package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login [hostname]",
	Short: "Obtain and save credentials for a remote host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	rootCmd.AddCommand(loginCmd)

	carapace.Gen(loginCmd).PositionalCompletion(
		net.ActionHosts(),
	)
}
