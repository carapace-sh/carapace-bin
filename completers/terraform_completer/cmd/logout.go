package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout [hostname]",
	Short: "Remove locally-stored credentials for a remote host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logoutCmd).Standalone()

	rootCmd.AddCommand(logoutCmd)

	carapace.Gen(logoutCmd).PositionalCompletion(
		net.ActionHosts(), // TODO complete authorized hosts
	)
}
