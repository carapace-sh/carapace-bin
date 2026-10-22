package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var server_unpauseCmd = &cobra.Command{
	Use:   "unpause",
	Short: "Unpause server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_unpauseCmd).Standalone()

	serverCmd.AddCommand(server_unpauseCmd)

	carapace.Gen(server_unpauseCmd).PositionalCompletion(action.ActionServers())
}
